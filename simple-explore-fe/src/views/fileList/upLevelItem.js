import React from "react";
import {ListItem, makeStyles} from "@material-ui/core";
import SubdirectoryArrowLeftIcon from "@material-ui/icons/SubdirectoryArrowLeft";
//import {useNavigate} from 'react-router-dom'

const useStyles = makeStyles((theme) => ({
  backIcon: {
    frontSize: 'large',
    htmlColor: '#64b5f6'
  }
}));

function UpLevelItem({listener}) {
  //const navigate = useNavigate()
  return (
    <ListItem button key={'up-level-item'} onClick={listener}>
      <SubdirectoryArrowLeftIcon className={useStyles.backIcon}/>
      <span>../</span>
    </ListItem>
  )
}

export default UpLevelItem;