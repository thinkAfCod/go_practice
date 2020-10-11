import 'react-perfect-scrollbar/dist/css/styles.css';
import 'video.js/dist/video-js.css'
import 'video-react/dist/video-react.css'
import React from 'react';
import {useRoutes} from 'react-router-dom';
import {ThemeProvider} from '@material-ui/core';
import GlobalStyles from 'src/components/GlobalStyles';
import 'src/mixins/chartjs';
import theme from 'src/theme';
import routes from 'src/routes';
// import axios from './utils/request';
//
// React.$axios = axios

function App() {
  const routing = useRoutes(routes)
  //debugger
  return (
    <ThemeProvider theme={theme}>
      <GlobalStyles/>
      {routing}
    </ThemeProvider>
  );
};

export default App;
