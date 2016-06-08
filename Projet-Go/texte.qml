import QtQuick 2.0
import QtQuick.Controls 1.1

Rectangle {
    id: root
    color: "light blue"

    width: 1200
    height: 500
TextEdit {
    id: textEdit1
    x: 8
    y: 8
    width: 304
    height: 20
    text: fileOp.content
    font.pixelSize: 18
	}
}
