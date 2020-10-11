import React from 'react';
import Page from 'src/components/Page';
import { makeStyles, Box, GridList, GridListTile } from '@material-ui/core';
import Crumb from '../../components/CrumbView';

const crumbClickHandler = (param) => {
  console.log('crumbClickHandler', param);
};

const crumbList = [
  {href: '/', handler: crumbClickHandler, title: '面包屑1'},
  {href: '/', handler: crumbClickHandler, title: '面包屑2'},
];

const exploreItem = [
  {
    id: 1,
    tabId: 1,
    parentItemId: 0,
    name: 'test',
    cover: '/test_file',
    uri: '/fdsafds',
    description: 'fdasf',
    keyword: 'fdsa'
  },
  {
    id: 2,
    tabId: 1,
    parentItemId: 0,
    name: 'test',
    cover: '/test_file',
    uri: '/fdsafds',
    description: 'fdasf',
    keyword: 'fdsa'
  },
  {
    id: 3,
    tabId: 1,
    parentItemId: 0,
    name: 'test',
    cover: '/test_file',
    uri: '/fdsafds',
    description: 'fdasf',
    keyword: 'fdsa'
  },
  {
    id: 4,
    tabId: 1,
    parentItemId: 0,
    name: 'test',
    cover: '/test_file',
    uri: '/fdsafds',
    description: 'fdasf',
    keyword: 'fdsa'
  },
  {
    id: 5,
    tabId: 1,
    parentItemId: 0,
    name: 'test',
    cover: '/test_file',
    uri: '/fdsafds',
    description: 'fdasf',
    keyword: 'fdsa'
  },
];

const useStyles = makeStyles((theme) => ({
  root: {
    backgroundColor: theme.palette.background.dark,
    minHeight: '100%',
    padding: theme.spacing(1),
  },
  container: {
    minHeight: '100%',
    minWidth: '100%'
  },
  gridList: {
    minHeight: '100%',
    minWidth: '100%',
  }
}));

const PreviewList = () => {
  const classes = useStyles();
  // todo 前面是面包屑，右边抵头是展示方式选项，列表、带封面两种模式
  return (
    <Page
      className={classes.root}
      title="Account"
    >
      <Box flexGrow={1}>
        <Crumb className="inherits" linkList={crumbList} handler={crumbClickHandler} />
      </Box>
      <GridList cellHeight={500} className={classes.gridList} cols={3}>
        {
          exploreItem.map((item) => {
            return (
              <GridListTile key={item.id} cols={1} onClick={()=> console.log('GridListTile has been clicked')}>
                <img src="https://th.bing.com/th/id/OIP.tEUHSnG-H2fQP_XdUSzKTQHaHa?pid=Api&rs=1" alt="???" />
              </GridListTile>
            );
          })
        }
      </GridList>
    </Page>
  );
};

export default PreviewList;
