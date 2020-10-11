import React from 'react';
import ReactDOM from 'react-dom';
import {makeStyles, CircularProgress} from '@material-ui/core';

// const useStyles = makeStyles((theme) => ({
//   loading_div: {
//     width: '100%',
//     height: '100%',
//     position: 'fixed',
//     textAlign: 'center',
//     zIndex: 999,
//   },
//   progress: {
//     position: 'absolute',
//     left: '50%',
//     top: '50%',
//     //transform: '-50% -50%',
//   }
// }))

export default class Loading extends React.Component {
  constructor(props) {
    super(props);
    console.log('loading component', props)
  }

  render() {
    return (
      <div>
        <CircularProgress />
      </div>
    )
  }
}

Loading.show = function() {
  this.div = document.createElement('div');
  document.body.appendChild(this.div);
  console.log(this.div)
  ReactDOM.render((<Loading {...this.props}/>), this.div);
}

Loading.remove = function() {
  this.div && ReactDOM.unmountComponentAtNode(this.div); //从div中移除已挂载的Loading组件
  this.div && this.div.parentNode.removeChild(this.div); //移除挂载的容器
}