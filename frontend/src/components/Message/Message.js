import React from 'react';
import "./Message.scss";

function Message(props){

  let msgData = JSON.parse(props.message);
  return <div className="Message">{msgData.body}</div>;
}

export default Message;