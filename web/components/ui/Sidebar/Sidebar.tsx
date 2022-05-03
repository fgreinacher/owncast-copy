import Sider from 'antd/lib/layout/Sider';
import { useRecoilValue } from 'recoil';
import { useEffect } from 'react';
import { ChatMessage } from '../../../interfaces/chat-message.model';
import ChatContainer from '../../chat/ChatContainer';
import {
  chatMessagesAtom,
  chatVisibilityAtom,
  chatStateAtom,
} from '../../stores/ClientConfigStore';
import { ChatState, ChatVisibilityState } from '../../../interfaces/application-state';
import ChatTextField from '../../chat/ChatTextField/ChatTextField';

export default function Sidebar() {
  const messages = useRecoilValue<ChatMessage[]>(chatMessagesAtom);
  const chatVisibility = useRecoilValue<ChatVisibilityState>(chatVisibilityAtom);
  const chatState = useRecoilValue<ChatState>(chatStateAtom);

  return (
    <Sider
      collapsed={chatVisibility === ChatVisibilityState.Hidden}
      collapsedWidth={0}
      width="100%"
    >
      <ChatContainer messages={messages} state={chatState} />
      <ChatTextField />
    </Sider>
  );
}
