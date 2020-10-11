import React, { useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import PropTypes from 'prop-types';
import {
  Box,
  Drawer,
  Hidden,
  List,
  makeStyles
} from '@material-ui/core';
import NavItem from './NavItem';

const tabItem = [
  {id: 0, tabName: "全部文件", tabParentId: 0, icon: "file", description: "", remark: ""},
  {id: 1, tabName: "视频", tabParentId: 0, icon: "movie", description: "", remark: ""},
  {id: 2, tabName: "写真", tabParentId: 0, icon: "camera", description: "", remark: ""},
  {id: 3, tabName: "音声", tabParentId: 0, icon: "album", description: "", remark: ""},
  {id: 4, tabName: "漫画", tabParentId: 0, icon: "image", description: "", remark: ""},
  {id: 5, tabName: "小说", tabParentId: 0, icon: "book", description: "", remark: ""},
];

const items = tabItem.map((tab) => {
  tab.path = tab.icon === 'movie' || tab.icon === 'album' || tab.icon === 'file' ? tab.icon : undefined;
  return {
    href: tab.path ? tab.path : '/',
    icon: tab.icon,
    title: tab.tabName
  };
});

const useStyles = makeStyles(() => ({
  mobileDrawer: {
    width: 256
  },
  desktopDrawer: {
    width: 256,
    top: 64,
    height: 'calc(100% - 64px)'
  },
  avatar: {
    cursor: 'pointer',
    width: 64,
    height: 64
  }
}));

const NavBar = ({ onMobileClose, openMobile }) => {
  const classes = useStyles();
  const location = useLocation();

  useEffect(() => {
    if (openMobile && onMobileClose) {
      onMobileClose();
    }
  }, [location.pathname]);

  const content = (
    <Box
      height="100%"
      display="flex"
      flexDirection="column"
    >
      <Box p={2}>
        <List>
          {items.map((item) => (
            <NavItem
              href={item.href}
              key={item.title}
              title={item.title}
              icon={item.icon}
            />
          ))}
        </List>
      </Box>
      <Box flexGrow={1} />
    </Box>
  );

  return (
    <>
      <Hidden lgUp>
        <Drawer
          anchor="left"
          classes={{ paper: classes.mobileDrawer }}
          onClose={onMobileClose}
          open={openMobile}
          variant="temporary"
        >
          {content}
        </Drawer>
      </Hidden>
      <Hidden mdDown>
        <Drawer
          anchor="left"
          classes={{ paper: classes.desktopDrawer }}
          open
          variant="persistent"
        >
          {content}
        </Drawer>
      </Hidden>
    </>
  );
};

NavBar.propTypes = {
  onMobileClose: PropTypes.func,
  openMobile: PropTypes.bool
};

NavBar.defaultProps = {
  onMobileClose: () => {
  },
  openMobile: false
};

export default NavBar;
