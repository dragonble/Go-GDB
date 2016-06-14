
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

		Rectangle {
		id: lineColumn
		property int rowHeight: textarea.font.pixelSize + 3
		color: "#f2f2f2"
		width: 50
		height: parent.height
		Rectangle {
		height: parent.height
		anchors.right: parent.right
		width: 1
		color: "#ddd"
		}
		Column {
		y: -textarea.flickableItem.contentY + 4
		width: parent.width
		Repeater {
		model: Math.max(textarea.lineCount + 2, (lineColumn.height/lineColumn.rowHeight) )
		delegate: Text {
		id: text
		state : 'normal'
		font: textarea.font
		width: lineColumn.width
		horizontalAlignment: Text.AlignHCenter
		verticalAlignment: Text.AlignVCenter
		height: lineColumn.rowHeight
		renderType: Text.NativeRendering
		text: index + 1	 
		MouseArea {
   		      anchors.fill: parent
   		      onClicked: {
		if (text.state == 'normal'){
		text.state = 'modif'
		fileOp.addbreakpoint(index+1)
		}
		else {
		text.state = 'normal'
		fileOp.rmvbreakpoint(index+1)
		}
		}
		     }
 
   		  states: [
      		   State {
      		       name: "modif"
      		       PropertyChanges { target: text; color : "blue" }
     		    },
			State {
    		         name: "normal"
      		       PropertyChanges { target: text; color : "#666" }
     		    }
   		  ]
		}
		}
		}
		}
		TextArea {
		id: textarea
		anchors.left: lineColumn.right
		anchors.right: parent.right
		anchors.top: parent.top 
		anchors.bottom: parent.bottom
		wrapMode: TextEdit.NoWrap
		frameVisible: false
		text: fileOp.content
		font.pixelSize: 18
		}
		}
	
            Rectangle {
                id: row2
                color: "lightgray"

		ToolBar{
		RowLayout {
            anchors.fill: parent
		ToolButton {
				iconSource: "Ressources2/back.png"
				onClicked : { fileOp.debugreverse() }
						
			}
		ToolButton {

				iconSource: "Ressources2/run.png"
				onClicked : { fileOp.debugrun()	}
			}
			
		ToolButton {

				iconSource: "Ressources2/step.png"
				onClicked : { fileOp.debugstep() }
			}
		ToolButton {
			
				iconSource: "Ressources2/continue.png"	
				onClicked : { fileOp.debugcontinue() }	
			}
		
			
			Item { Layout.fillWidth: true }
 
		
		}
		
	}
            }
        }

    }
	
	
}