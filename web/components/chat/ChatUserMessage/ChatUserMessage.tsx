/* eslint-disable react/no-danger */
import { useEffect, useState } from 'react';
import { Highlight } from 'react-highlighter-ts';
import he from 'he';
import cn from 'classnames';
import s from './ChatUserMessage.module.scss';
import { formatTimestamp } from './messageFmt';
import { ChatMessage } from '../../../interfaces/chat-message.model';
import { ModIcon } from '../../ui';
import ChatModerationActionMenu from '../ChatModerationActionMenu/ChatModerationActionMenu';

interface Props {
  message: ChatMessage;
  showModeratorMenu: boolean;
  highlightString: string;
  sentBySelf: boolean;
  sameUserAsLast: boolean;
  isAuthorModerator: boolean;
}

export default function ChatUserMessage({
  message,
  highlightString,
  showModeratorMenu,
  sentBySelf, // Move the border to the right and render a background
  sameUserAsLast,
  isAuthorModerator,
}: Props) {
  const { id: messageId, body, user, timestamp } = message;
  const { id: userId, displayName, displayColor } = user;

  const color = `var(--theme-user-colors-${displayColor})`;
  const formattedTimestamp = `Sent at ${formatTimestamp(timestamp)}`;
  const [formattedMessage, setFormattedMessage] = useState<string>(body);

  useEffect(() => {
    setFormattedMessage(he.decode(body));
  }, [message]);

  return (
    <div style={{ padding: 3.5 }}>
      <div
        className={cn(s.root, {
          [s.ownMessage]: sentBySelf,
        })}
        style={{ borderColor: color }}
        title={formattedTimestamp}
      >
        {!sameUserAsLast && (
          <div className={s.user} style={{ color }}>
            <span className={s.userName}>{displayName}</span>
            {isAuthorModerator && <ModIcon />}
          </div>
        )}
        <Highlight search={highlightString}>
          <div className={s.message}>{formattedMessage}</div>
        </Highlight>
        {showModeratorMenu && (
          <div className={s.modMenuWrapper}>
            <ChatModerationActionMenu
              messageID={messageId}
              accessToken=""
              userID={userId}
              userDisplayName={displayName}
            />
          </div>
        )}
        <div className={s.customBorder} style={{ color }} />
        <div className={s.background} style={{ color }} />
      </div>
    </div>
  );
}
