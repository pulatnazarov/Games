from PyQt5.QtWidgets import QApplication, QWidget, QPushButton, QRadioButton, QLabel, QMessageBox, QHBoxLayout
from PyQt5.QtGui import QFont
import sys

class Button(QPushButton):
    def __init__(self, name, oyna, x, y):
        super().__init__(name, oyna)
        self.setFont(QFont('Times', 30))
        self.setGeometry(x, y, 100, 100)
    def click(self, func):
        self.clicked.connect(func)

class Window(QWidget):
    def __init__(self):
        super().__init__()
        self.setWindowTitle('X & O (2 players)')
        self.setGeometry(1000, 300, 340, 340)
        self.control()
    def font(self, ob):
        ob.setFont(QFont('Times', 30))
    def fon(self, obj):
        obj.setFont(QFont('Times', 20))
    def control(self):
        self.choice=QLabel('Choose', self)
        self.choice.move(130, 10)
        self.fon(self.choice)
        self.v1=QRadioButton('X', self)
        self.v1.move(100, 40)
        self.v2=QRadioButton('O', self)
        self.v2.move(200, 40)

        self.ok=Button('ok', self, 160, 160)
        hbox=QHBoxLayout(self)
        hbox.addStretch()
        hbox.addWidget(self.ok)
        hbox.addStretch()

        self.new1=[self.v2, self.v1, self.ok, self.choice]

        self.a1=Button('', self, 10, 10)
        self.a2=Button('', self, 120, 10)
        self.a3=Button('', self, 230, 10)
        self.a4=Button('', self, 10, 120)
        self.a5=Button('', self, 120, 120)
        self.a6=Button('', self, 230, 120)
        self.a7=Button('', self, 10, 230)
        self.a8=Button('', self, 120, 230)
        self.a9=Button('', self, 230, 230)
        self.new=[self.a1, self.a2, self.a3, self.a4, self.a5, self.a6, self.a7, self.a8, self.a9]
        
        for item in self.new:
            item.hide()
        self.ok.click(self.run)

        self.a1.click(self.A1)
        self.a2.click(self.A2)
        self.a3.click(self.A3)
        self.a4.click(self.A4)
        self.a5.click(self.A5)
        self.a6.click(self.A6)
        self.a7.click(self.A7)
        self.a8.click(self.A8)
        self.a9.click(self.A9)

    def run(self):
        if self.v1.isChecked():
            self.x='X'
        elif self.v2.isChecked():
            self.x='O'
        else:
            self.x='X'
        for item in self.new:
            item.show()
        for item in self.new1:
            item.hide()
    
    def scan(self):
        win=QMessageBox(self)
        win.setWindowTitle("X & O")

        if self.a1.text()!='' and self.a1.text()==self.a2.text() and self.a1.text()==self.a3.text():
            win.setText(self.a1.text()+' WINS!!!')
            win.show()
            self.refresh()
        
        elif self.a1.text()!='' and self.a1.text()==self.a4.text() and self.a1.text()==self.a7.text():
            win.setText(self.a1.text()+' WINS!!!')
            win.show()
            self.refresh()
        
        elif self.a1.text()!='' and self.a1.text()==self.a5.text() and self.a1.text()==self.a9.text():
            win.setText(self.a1.text()+' WINS!!!')
            win.show()
            self.refresh()
        
        elif self.a4.text()!='' and self.a4.text()==self.a5.text() and self.a5.text()==self.a6.text():
            win.setText(self.a4.text()+' WINS!!!')
            win.show()
            self.refresh()
        
        elif self.a8.text()!='' and self.a8.text()==self.a9.text() and self.a9.text()==self.a7.text():
            win.setText(self.a8.text()+' WINS!!!')
            win.show()
            self.refresh()
        
        elif self.a3.text()!='' and self.a5.text()==self.a7.text() and self.a5.text()==self.a3.text():
            win.setText(self.a3.text()+' WINS!!!')
            win.show()
            self.refresh()
        
        elif self.a3.text()!='' and self.a3.text()==self.a6.text() and self.a3.text()==self.a9.text():
            win.setText(self.a1.text()+' WINS!!!')
            win.show()
            self.refresh()
        
        elif self.a2.text()!='' and self.a2.text()==self.a5.text() and self.a2.text()==self.a8.text():
            win.setText(self.a2.text()+' WINS!!!')
            win.show()
            self.refresh()
        
        elif self.a1.text()!='' and self.a2.text()!='' and self.a3.text()!='' and self.a4.text()!='' and self.a5.text()!='' and self.a6.text()!='' and self.a7.text()!='' and self.a8.text()!='' and self.a9.text()!='':
            win.setText('DRAW!!!')
            win.show()
            self.refresh()

    def refresh(self):
        for item in self.new:
            item.setText("")
            item.hide()
        for item in self.new1:
            item.show()

    def A1(self):
        if self.a1.text()=='':
            self.a1.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    
    def A2(self):
        if self.a2.text()=='':
            self.a2.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    
    def A3(self):
        if self.a3.text()=='':
            self.a3.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    
    def A4(self):
        if self.a4.text()=='':
            self.a4.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    
    def A5(self):
        if self.a5.text()=='':
            self.a5.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    
    def A6(self):
        if self.a6.text()=='':
            self.a6.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    
    def A7(self):
        if self.a7.text()=='':
            self.a7.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    
    def A8(self):
        if self.a8.text()=='':
            self.a8.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    
    def A9(self):
        if self.a9.text()=='':
            self.a9.setText(self.x)
            if self.x=='X':
                self.x='O'
            else:
                self.x='X'
        self.scan()
    

app=QApplication(sys.argv)
obj=Window()
obj.show()
sys.exit(app.exec_())