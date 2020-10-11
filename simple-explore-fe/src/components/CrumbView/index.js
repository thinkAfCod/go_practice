import React from 'react';
import { Breadcrumbs, Link } from '@material-ui/core';
import PropTypes from 'prop-types';

const Crumb = ({ className, linkList, handler }) => {
  return (
    <Breadcrumbs separator="/">
      {
        linkList.map((item) => (
          <Link key={item.title} href="#" color="inherit" onClick={() => handler(item.title)}>{item.title}</Link>))
      }
    </Breadcrumbs>
  );
};

Crumb.propTypes = {
  className: PropTypes.string,
  linkList: PropTypes.array,
  handler: PropTypes.func
};

export default Crumb;
