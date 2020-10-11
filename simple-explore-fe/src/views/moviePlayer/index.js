import React from "react";
import VideoJsPlayer from "./VideoJsPlayer";
import {Container} from "@material-ui/core";
import {useParams} from 'react-router-dom'

function MoviePlayer(props) {
  let params = useParams()
  console.log()
  const videoJsOptions = {
    autoplay: false,
    controls: true,
    preload: 'auto',
    sources: [{
      src: `http://192.168.50.109:8080/file?id=${params.fileId}`,
      type: params.mediaType
    }],
    fluid: true,
  }
  console.log('MoviePlayer', videoJsOptions)
  return (
    <Container>
      <VideoJsPlayer {...videoJsOptions} />
    </Container>
  )
}


export default MoviePlayer;