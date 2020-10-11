import React from 'react';

const Logo = (props) => {
  return (
    <img
      alt="Logo"
      src="/static/favicon.ico"
      {...props}
    />
  );
};

export default Logo;
