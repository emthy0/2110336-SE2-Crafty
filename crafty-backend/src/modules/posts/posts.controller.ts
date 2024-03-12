import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
} from '@nestjs/common'
import { PostsService } from './posts.service'
import { CreatePostDto } from './dto/create-post.dto'
import { UpdatePostDto } from './dto/update-post.dto'
import { ApiTags } from '@nestjs/swagger'
import { CreateReviewDto } from './dto/create-review'
import { AddFavoriteDto } from './dto/add-favorite.dto'

@ApiTags('posts')
@Controller('posts')
export class PostsController {
  constructor(private readonly postsService: PostsService) {}

  @Post(':id/reviews')
  addReview(
    @Param('id') postId: string,
    @Body() createReviewDto: CreateReviewDto,
  ) {
    return this.postsService.addReview(
      postId,
      createReviewDto.desc,
      createReviewDto.rate,
      createReviewDto.sender,
    )
  }

  @Get(':id/reviews')
  getReviews(@Param('id') postId: string) {
    return this.postsService.getReviews(postId)
  }

  @Post(':id/favorites')
  addFavorite(
    @Param('id') postId : string,
    @Body() addFavoriteDto: addFavoriteDto,
  ){
    return this.postsService.get
  }

  @Post()
  create(@Body() createPostDto: CreatePostDto) {
    return this.postsService.create(createPostDto)
  }

  @Get()
  findAll() {
    return this.postsService.findAll()
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.postsService.findOne(id)
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updatePostDto: UpdatePostDto) {
    return this.postsService.update(id, updatePostDto)
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.postsService.remove(id)
  }
}
