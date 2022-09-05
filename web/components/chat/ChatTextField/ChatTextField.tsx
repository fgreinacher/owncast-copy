import { SendOutlined, SmileOutlined } from '@ant-design/icons';
import { Button, Popover } from 'antd';
import React, { FC, useState } from 'react';
import { useRecoilValue } from 'recoil';
import { Transforms, createEditor, BaseEditor, Text } from 'slate';
import { Slate, Editable, withReact, ReactEditor } from 'slate-react';
import { EmojiPicker } from './EmojiPicker';
import WebsocketService from '../../../services/websocket-service';
import { websocketServiceAtom } from '../../stores/ClientConfigStore';
import { MessageType } from '../../../interfaces/socket-events';
import styles from './ChatTextField.module.scss';

type CustomElement = { type: 'paragraph' | 'span'; children: CustomText[] };
type CustomText = { text: string };

declare module 'slate' {
  interface CustomTypes {
    Editor: BaseEditor & ReactEditor;
    Element: CustomElement;
    Text: CustomText;
  }
}

// eslint-disable-next-line react/prop-types
const Image = ({ element }) => (
  <img
    // eslint-disable-next-line no-undef
    // eslint-disable-next-line react/prop-types
    src={element.url}
    alt="emoji"
    style={{ display: 'inline', position: 'relative', width: '30px', bottom: '10px' }}
  />
);

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const insertImage = (editor, url) => {
  // const text = { text: '' };
  // const image: ImageElement = { type: 'image', url, children: [text] };
  // Transforms.insertNodes(editor, image);
};

const withImages = editor => {
  const { isVoid } = editor;

  // eslint-disable-next-line no-param-reassign
  editor.isVoid = element => (element.type === 'image' ? true : isVoid(element));
  // eslint-disable-next-line no-param-reassign
  editor.isInline = element => element.type === 'image';

  return editor;
};

export type EmptyText = {
  text: string;
};

// type ImageElement = {
//   type: 'image';
//   url: string;
//   children: EmptyText[];
// };

const Element = (props: any) => {
  const { attributes, children, element } = props;

  switch (element.type) {
    case 'image':
      return <Image {...props} />;
    default:
      return <p {...attributes}>{children}</p>;
  }
};

const serialize = node => {
  if (Text.isText(node)) {
    const string = node.text;
    // if (node.bold) {
    //   string = `<strong>${string}</strong>`;
    // }
    return string;
  }

  const children = node.children.map(n => serialize(n)).join('');

  switch (node.type) {
    case 'paragraph':
      return `<p>${children}</p>`;
    case 'image':
      return `<img src="${node.url}" alt="emoji" />`;
    default:
      return children;
  }
};

export type ChatTextFieldProps = {};

export const ChatTextField: FC<ChatTextFieldProps> = () => {
  const [showEmojis, setShowEmojis] = useState(false);
  const websocketService = useRecoilValue<WebsocketService>(websocketServiceAtom);
  const [editor] = useState(() => withImages(withReact(createEditor())));

  const sendMessage = () => {
    if (!websocketService) {
      console.log('websocketService is not defined');
      return;
    }

    const message = serialize(editor);

    websocketService.send({ type: MessageType.CHAT, body: message });

    // Clear the editor.
    Transforms.select(editor, [0, editor.children.length - 1]);
    Transforms.delete(editor);
  };

  const handleChange = () => {};

  const handleEmojiSelect = (e: any) => {
    ReactEditor.focus(editor);

    if (e.url) {
      // Custom emoji
      const { url } = e;
      insertImage(editor, url);
    } else {
      // Native emoji
      const { emoji } = e;
      Transforms.insertText(editor, emoji);
    }
  };

  const onKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      sendMessage();
    }
  };

  return (
    <div>
      <div className={styles.root}>
        <Slate
          editor={editor}
          value={[{ type: 'paragraph', children: [{ text: '' }] }]}
          onChange={handleChange}
        >
          <Editable
            onKeyDown={onKeyDown}
            renderElement={p => <Element {...p} />}
            placeholder="Chat message goes here..."
            style={{ width: '100%' }}
            autoFocus
          />
          <Popover
            content={<EmojiPicker onEmojiSelect={handleEmojiSelect} />}
            trigger="click"
            onVisibleChange={visible => setShowEmojis(visible)}
            visible={showEmojis}
          />
        </Slate>

        <button
          type="button"
          className={styles.emojiButton}
          title="Emoji picker button"
          onClick={() => setShowEmojis(!showEmojis)}
        >
          <SmileOutlined />
        </button>
        <Button
          className={styles.sendButton}
          size="large"
          type="ghost"
          icon={<SendOutlined />}
          onClick={sendMessage}
        />
      </div>
    </div>
  );
};
