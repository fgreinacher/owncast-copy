import Sider from 'antd/lib/layout/Sider';
import { useRecoilValue } from 'recoil';
import { FC } from 'react';
import { ChatMessage } from '../../../interfaces/chat-message.model';
import { ChatContainer } from '../../chat/ChatContainer/ChatContainer';
import s from './Sidebar.module.scss';

import {
  chatDisplayNameAtom,
  chatUserIdAtom,
  isChatModeratorAtom,
  visibleChatMessagesSelector,
} from '../../stores/ClientConfigStore';

export const Sidebar: FC = () => {
  const chatDisplayName = useRecoilValue<string>(chatDisplayNameAtom);
  const chatUserId = useRecoilValue<string>(chatUserIdAtom);
  const isChatModerator = useRecoilValue<boolean>(isChatModeratorAtom);
  const messages = useRecoilValue<ChatMessage[]>(visibleChatMessagesSelector);

  return (
    <Sider className={s.root} collapsedWidth={0} width={320}>
      <ChatContainer
        messages={messages}
        usernameToHighlight={chatDisplayName}
        chatUserId={chatUserId}
        isModerator={isChatModerator}
      />
    </Sider>
  );
};
export default Sidebar;
