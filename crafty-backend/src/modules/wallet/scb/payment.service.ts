import { Injectable } from '@nestjs/common'
import axios from 'axios'
import { v4 as uuidv4 } from 'uuid'

interface QRResponse {
  qrImage: string

  ref: string
  tid: string
}

@Injectable()
export class PaymentService {
  constructor() {
    // this.
  }
  private generateRandomString(length: number): string {
    const charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
    let result = ''
    for (let i = 0; i < length; i++) {
      const randomIndex = Math.floor(Math.random() * charset.length)
      result += charset[randomIndex]
    }
    return result
  }

  private async generateOAuthToken(): Promise<string> {
    const apiKey = process.env.SCB_API_KEY
    const apiSecret = process.env.SCB_API_SECRET
    const reqUUID = uuidv4()
    const header = {
      'Content-Type': 'application/json',
      resourceOwnerId: apiKey,
      requestUid: reqUUID,
      'accept-language': 'EN',
    }
    const data = {
      applicationKey: apiKey,
      applicationSecret: apiSecret,
    }

    const res = await axios.post(
      'https://api-sandbox.partners.scb/partners/sandbox/v1/oauth/token',
      data,
      { headers: header },
    )
    const status = res.data.status.code
    // consol
    if (status !== 1000) {
      throw new Error('Failed to generate OAuth token')
    }
    // console.log(res.data.data)
    const accessToken = res.data.data.accessToken

    return accessToken
  }

  async generateQR(amount: number): Promise<QRResponse> {
    const accessToken = await this.generateOAuthToken()
    const reqUUID = uuidv4()
    const apiKey = process.env.SCB_API_KEY
    const ref1 = this.generateRandomString(20)
    const header = {
      'content-type': 'application/json',
      resourceOwnerId: process.env.SCB_API_KEY,
      requestUid: reqUUID,
      'accept-language': 'EN',
      authorization: `Bearer ${accessToken}`,
    }
    const data = {
      qrType: 'PPCS',
      amount: amount,
      invoice: ref1,
      merchantId: process.env.SCB_MERCHANT_ID,
      terminalId: process.env.SCB_TERMINAL_ID,
      ppType: 'BILLERID',
      ppId: process.env.SCB_BILLER_ID,
      ref1: ref1,
      ref3: 'CRT888888',
    }
    console.log(data)
    console.log(header)
    const res = await axios
      .post(
        'https://api-sandbox.partners.scb/partners/sandbox/v1/payment/qrcode/create',
        data,
        { headers: header },
      )
      .catch((err) => {
        console.log(err.code)
        return err.response
      })
    console.log(res.data.status)
    if (res.data.status.code !== 1000) {
      throw new Error('Failed to generate QR code')
    }
    return {
      qrImage: res.data.data.qrImage,
      ref: ref1,
      tid: reqUUID,
    }
  }
}
