import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0

ApplicationWindow {
    visible: true
    width: 1200
    height: 600

    SplitView {
        anchors.fill: parent

        Rectangle {
            id: column
            width: 200
            Layout.minimumWidth: 100
            Layout.maximumWidth: 300
            color: "lightsteelblue"
        }

        SplitView {
            orientation: Qt.Vertical
            Layout.fillWidth: true

            Rectangle {
                id: row1
                height: 400
                color: "lightblue"
                Layout.minimumHeight: 1

		

		TextArea {
   		    id: textEdit1
   		     x: 8
   		     y: 8
   		     width: row1.width - 20
    		     height: row1.height - 20
    		     text: fileOp.content
   		     font.pixelSize: 18
		     wrapMode: TextEdit.Wrap
			clip : true

		        }
			 
		    
            	}
	
            Rectangle {
                id: row2
                color: "lightgray"
            }
        }
    }
}
