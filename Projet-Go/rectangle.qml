import QtQuick 2.0

Rectangle {
	id: screen
        width: 800
        height: 400
       

        Rectangle {
                width: screen.width; height: 250
		anchors.left: screen.left
		anchors.bottom: screen.bottom
                
                color: "red"

                MouseArea {
                        anchors.fill: parent
                        onClicked: console.log("Stop poking me!")
                }
        }
}
