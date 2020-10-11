import {useNavigate} from "react-router-dom";
import {ListItem} from "@material-ui/core";
import React from "react";
import SubjectIcon from "@material-ui/icons/Subject";
import FolderIcon from "@material-ui/icons/Folder";
import AlbumIcon from "@material-ui/icons/Album";
import PhotoIcon from "@material-ui/icons/Photo";
import MovieIcon from '@material-ui/icons/Movie';
import {makeStyles} from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  name: {
    marginLeft: '10px'
  }
}))

function FileItem({fileData, listener}) {
  const navigate = useNavigate()
  const classes = useStyles()
  let labelId = `file-item-view-${fileData.id}`
  let iconPath = getIconAndPath(fileData, classes)
  let isDir = fileData.mediaType === undefined || fileData.mediaType === ''
  const leaveTo = () => {
    navigate(iconPath.path)
    if (isDir) {
      listener(fileData.id)
    }
  }

  return (
    <ListItem button key={labelId} onClick={leaveTo}>
      {iconPath.icon}
      <span className={classes.name}>{fileData.name}</span>
    </ListItem>
  )
}

function getIconAndPath(fileData, classes) {
  let icon;
  let path;
  let mediaType = fileData.mediaType;
  let id = fileData.id;
  let fileId = fileData.fileId;
  let parentId = fileData.parentItemId;
  if (mediaType === undefined || mediaType === '') {
    icon = <FolderIcon style={{color: '#ffb74d'}}/>
    path = `/app/file/${id}`
  } else if (mediaType.includes('text')) {
    icon = <SubjectIcon/>
    path = `/app/text/${parentId}/${fileId}`
  } else if (mediaType.includes('image')) {
    icon = <PhotoIcon style={{color: '#4caf50'}}/>
    path = `/app/text/${parentId}/${fileId}`
  } else if (mediaType.includes('video')) {
    icon = <MovieIcon color={'secondary'}/>
    //movie/:parentId/:fileId/:mediaType
    path = `/app/movie/${parentId}/${fileId}/${encodeURIComponent(mediaType)}`
  } else if (mediaType.includes('audio')) {
    icon = <AlbumIcon color={'action'}/>
    path = `/app/audio/${parentId}/${fileId}`
  }
  return {
    path: path,
    icon: icon
  };
}

export default FileItem;