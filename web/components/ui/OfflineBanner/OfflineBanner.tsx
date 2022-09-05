import { Divider, Button } from 'antd';
import { NotificationFilled } from '@ant-design/icons';
import { FC } from 'react';
import styles from './OfflineBanner.module.scss';

export type OfflineBannerProps = {
  name: string;
  text: string;
};

export const OfflineBanner: FC<OfflineBannerProps> = ({ name, text }) => (
  <div className={styles.outerContainer}>
    <div className={styles.innerContainer}>
      <div className={styles.header}>{name} is currently offline.</div>
      <Divider />
      <div>{text}</div>

      <div className={styles.footer}>
        <Button type="primary" onClick={() => console.log('show notification modal')}>
          <NotificationFilled />
          Notify when Live
        </Button>
      </div>
    </div>
  </div>
);
