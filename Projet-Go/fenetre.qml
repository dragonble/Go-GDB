import QtQuick 2.2
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0

ApplicationWindow {
    visible: true
    width: 1000
    height: 600
	title: "Go-GDB"

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
                height: 200
                color: "lightblue"
                Layout.minimumHeight: 1
            }

            Rectangle {
                id: row2
                color: "lightgray"
            }
        }
    }
	ToolBar{
		RowLayout {
            anchors.fill: parent
		ToolButton {
				iconSource: "Ressources/back.png"	
				//onClicked: "Action"	
			}
		ToolButton {

				iconSource: "Ressources/run.png"	
				//onClicked: "Action"		
			}
			
		ToolButton {

				iconSource: "Ressources/step.png"
				//onClicked: "Action"	
			}
		ToolButton {
			
				iconSource: "Ressources/continue.png"
				//onClicked: "Action"			
			}

		ToolButton {
			
				iconSource: "Ressources/backtrace.png"	
				//onClicked: "Action"		
			}

		ToolButton {
			
				iconSource: "Ressources/variables.png"	
				//onClicked: "Action"	
	
			}

		
			
			Item { Layout.fillWidth: true }
            CheckBox {
                text: "Enabled"
                checked: true
                Layout.alignment: Qt.AlignRight
			}
		
		}
		
	}
}
