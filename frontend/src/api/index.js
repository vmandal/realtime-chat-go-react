
var socket = new WebSocket("ws://localhost:8000/ws");

let connect = cb => {
    console.log("connecting...");

    socket.onopen = () => {
        console.log("successfully connected");        
    }

    socket.onmessage = msg => {
        console.log('Received: ');
        console.log(msg);
        cb(msg); // how does this work?
    };

    socket.onclose = event => {
        console.log("socket closed.", event);
    }

    socket.onerror = error => {
        console.log("soecket error: ", error);
    }

};

let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg); // msg sent to Go listner at :8000
}

export {connect, sendMsg};