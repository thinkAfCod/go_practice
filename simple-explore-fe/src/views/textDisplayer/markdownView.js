import React, {useEffect, useState} from 'react';


const MarkdownView = function ({className, content}) {
  const screenResize = () => {
    console.log('window.innerHeight - 64', window.innerHeight - 64)
    return window.innerHeight - 64
  }
  const [height,setHeight] = useState(screenResize())
  //const [classes, setClasses] = React.useState({...className, height: screenResize()})
  const showdown = require('showdown')
  const conv = new showdown.Converter()
  const divRef = React.useRef()


  const testResize = React.useCallback(() => {
    debugger
    setHeight(screenResize())
  }, [screenResize])

  useEffect(() => {
    window.addEventListener('resize', testResize)
    return () => {
      window.removeEventListener('resize', testResize);
    }
  }, [testResize])

  return (
    <div ref={divRef} className={className} id="test-mark-down-1" style={{height: height}}
         dangerouslySetInnerHTML={{__html: conv.makeHtml(content)}}>
    </div>
  )
}

export default MarkdownView;