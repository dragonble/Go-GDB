 import QtQuick 2.0
 
 Rectangle {
     id: myRect
     width: 500; height: 500
     state : 'normal'
 
     MouseArea {
         anchors.fill: parent
         onClicked: {
if (myRect.state == 'normal')
myRect.state = 'modif'
else 
myRect.state = 'normal'
}
     }
 
     states: [
         State {
             name: "modif"
             PropertyChanges { target: myRect; color : "blue" }
         },
	State {
             name: "normal"
             PropertyChanges { target: myRect; color : "red" }
         }
     ]
 }
