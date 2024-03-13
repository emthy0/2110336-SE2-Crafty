import axios from 'axios';
import React, { useEffect, useState } from 'react';
import SidebarBubble from './_component/SidebarBubble';
import { apiPath, getMyName, getMyToken } from '@/app/_common/utils/chatting';

type ChatLayoutProps = {
  children: React.ReactNode;
};

const token = getMyToken();

const getAllChatrooms = async (): Promise<ReadChatroom[]> => {
  // using axios to fetch from localhost:5000/chats
  // add bearer token to the header

  // const token = localStorage.getItem('token');
  const config = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };

  const res = await axios.get<ReadChatroom[]>(`${apiPath}/chats`, config);
  return res.data;
};

const getSidebarData = async (): Promise<SidebarData[]> => {
  console.log('Sidebar reload!');

  const chatrooms = await getAllChatrooms();
  const myName = await getMyName(token);

  const talkersWithLastChatTime = chatrooms.map((chatroom: ReadChatroom) => {
    return {
      talkerName:
        chatroom.crafter.username === myName
          ? chatroom.craftee.username
          : chatroom.crafter.username,
      lastChatTime: chatroom.lastChatTime,
      chatroomId: chatroom.id,
    };
  });
  // console.log(talkersWithLastChatTime);
  return talkersWithLastChatTime;
};

const ChatLayout: React.FC<ChatLayoutProps> = async ({ children }) => {
  const sidebarDatas = await getSidebarData();

  return (
    <div className="flex h-screen">
      <aside className="flex w-48 flex-col gap-2 bg-ct_brown-100 pt-4" aria-label="Sidebar">
        <div className="mb-2 flex justify-center border-b pb-4">Chat History</div>
        {sidebarDatas.map((sidebarData: SidebarData) => {
          return (
            <SidebarBubble
              key={sidebarData.chatroomId}
              talkerName={sidebarData.talkerName}
              lastChatTime={sidebarData.lastChatTime}
              chatroomId={sidebarData.chatroomId}
            />
          );
        })}
      </aside>
      <section className="flex-1">
        {children} {/* This is your main content area */}
      </section>
    </div>
  );
};

export default ChatLayout;
