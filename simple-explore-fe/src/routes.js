import React from 'react';
import {Navigate} from 'react-router-dom';
import DashboardLayout from 'src/layouts/DashboardLayout';
import MainLayout from 'src/layouts/MainLayout';
import NotFoundView from 'src/views/errors/NotFoundView';
import PreviewList from 'src/views/previewList/PreviewList';
import MoviePlayer from "./views/moviePlayer";
import FileList from "./views/fileList";
import TextDisplayer from "./views/textDisplayer";
import PictureViewer from "./views/pictureViewer";

const routes = [
  {
    path: 'app',
    element: <DashboardLayout/>,
    children: [
      {path: 'movie/:parentId/:fileId/:mediaType', element: <MoviePlayer/>},
      {path: 'file/*', element: <FileList/>},
      {path: 'text/:parentId/:fileId', element: <TextDisplayer/>},
      {path: 'pic/:parentId/:index', element: <PictureViewer/>},
      {path: 'album', element: <PreviewList/>},
      {path: '*', element: <Navigate to="/404"/>}
    ]
  },
  {
    path: '/',
    element: <MainLayout/>,
    children: [
      {path: '404', element: <NotFoundView/>},
      {path: '/', element: <Navigate to="/app/pic/0/1"/>},
      {path: '*', element: <Navigate to="/404"/>}
    ]
  }
];

export default routes;
