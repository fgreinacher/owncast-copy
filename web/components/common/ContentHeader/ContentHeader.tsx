import cn from 'classnames';
import { FC } from 'react';
import { ServerLogo } from '../../ui';
import SocialLinks from '../../ui/SocialLinks/SocialLinks';
import { SocialLink } from '../../../interfaces/social-link.model';
import styles from './ContentHeader.module.scss';

export type ContentHeaderProps = {
  name: string;
  title: string;
  summary: string;
  tags: string[];
  links: SocialLink[];
  logo: string;
};

export const ContentHeader: FC<ContentHeaderProps> = ({
  name,
  title,
  summary,
  logo,
  tags,
  links,
}) => (
  <div className={styles.root}>
    <div className={styles.logoTitleSection}>
      <div className={styles.logo}>
        <ServerLogo src={logo} />
      </div>
      <div className={styles.titleSection}>
        <div className={cn(styles.title, styles.row)}>{name}</div>
        <div className={cn(styles.subtitle, styles.row)}>{title || summary}</div>
        <div className={cn(styles.tagList, styles.row)}>
          {tags.length > 0 && tags.map(tag => <span key={tag}>#{tag}&nbsp;</span>)}
        </div>
        <div className={cn(styles.socialLinks, styles.row)}>
          <SocialLinks links={links} />
        </div>
      </div>
    </div>
  </div>
);

export default ContentHeader;
