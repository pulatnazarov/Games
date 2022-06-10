
from PyQt5.QtWidgets import QApplication, QWidget, QLabel, QPushButton
from PyQt5.QtGui import QFont
import sys
ls=list()

class Button(QPushButton):
    def __init__(self, name, oyna, x, y):
        super().__init__(name, oyna)
        if name=='⌫':
            self.setFont(QFont("Times", 25))
            self.setGeometry(x, y, 45, 35)
        elif name=='AC':
            self.setFont(QFont("Times", 30))
            self.setGeometry(x, y, 75, 160)
        elif name=='0' or name=='=':
            self.setFont(QFont("Times", 30))
            self.setGeometry(x, y, 160, 75)
        else:
            self.setFont(QFont("Times", 30))
            self.setGeometry(x, y, 75, 75)
    def click(self, func):
        self.clicked.connect(func)

class Window(QWidget):
    def __init__(self):
        super().__init__()
        self.setWindowTitle('Calculator')
        self.setGeometry(1300, 300, 350, 585)
        #self.setStyleSheet("background-color: cyan;")
        self.control()
    def font(self, ob):
        ob.setFont(QFont("Times", 20))
    def control(self):
        self.result=QLabel('', self)
        self.result.move(10, 40)
        self.font(self.result)

        self.seven=Button('7', self, 10, 415)
        self.four=Button('4', self, 10, 330)
        self.one=Button('1', self, 10, 245)
        self.eight=Button('8', self, 95, 415)
        self.five=Button('5', self, 95, 330)
        self.two=Button('2', self, 95, 245)
        self.nine=Button('9', self, 180, 415)
        self.six=Button('6', self, 180, 330)
        self.three=Button('3', self, 180, 245)
        self.zero=Button('0', self, 10, 500)

        self.plus=Button('+', self, 10, 160)
        self.minus=Button('-', self, 95, 160)
        self.multiple=Button('×', self, 180, 160)
        self.divide=Button('÷', self, 265, 160)
        self.ac=Button('AC', self, 265, 245)
        self.float=Button('.', self, 265, 415)
        self.c=Button('⌫', self, 295, 115)
        self.equal=Button('=', self, 180, 500)
  
        self.one.click(self.wone)
        self.two.click(self.wtwo)
        self.three.click(self.wthree)
        self.four.click(self.wfour)
        self.five.click(self.wfive)
        self.six.click(self.wsix)
        self.seven.click(self.wseven)
        self.eight.click(self.weight)
        self.nine.click(self.wnine)
        self.zero.click(self.wzero)
        
        self.plus.click(self.wplus)
        self.minus.click(self.wminus)
        self.multiple.click(self.wmultiple)
        self.divide.click(self.wdivide)
        self.float.click(self.wfloat)
        
        self.ac.click(self.clearall)
        self.c.click(self.clear)

        self.equal.click(self.answer)

    def wone(self):
        ls.append('1')
        self.result.setText(self.result.text()+'1')
        self.result.adjustSize()
    def wtwo(self):
        ls.append('2')
        self.result.setText(self.result.text()+'2')
        self.result.adjustSize()
    def wthree(self):
        ls.append('3')
        self.result.setText(self.result.text()+'3')
        self.result.adjustSize()
    def wfour(self):
        ls.append('4')
        self.result.setText(self.result.text()+'4')
        self.result.adjustSize()
    def wfive(self):
        ls.append('5')
        self.result.setText(self.result.text()+'5')
        self.result.adjustSize()
    def wsix(self):
        ls.append('6')
        self.result.setText(self.result.text()+'6')
        self.result.adjustSize()
    def wseven(self):
        ls.append('7')
        self.result.setText(self.result.text()+'7')
        self.result.adjustSize()
    def weight(self):
        ls.append('8')
        self.result.setText(self.result.text()+'8')
        self.result.adjustSize()
    def wnine(self):
        ls.append('9')
        self.result.setText(self.result.text()+'9')
        self.result.adjustSize()
    def wzero(self):
        ls.append('0')
        self.result.setText(self.result.text()+'0')
        self.result.adjustSize()
    
    def wplus(self):
        ls.append('+')
        self.result.setText(self.result.text()+'+')
        self.result.adjustSize()
    def wminus(self):
        ls.append('-')
        self.result.setText(self.result.text()+'-')
        self.result.adjustSize()
    def wmultiple(self):
        ls.append('*')
        self.result.setText(self.result.text()+'×')
        self.result.adjustSize()
    def wdivide(self):
        ls.append('/')
        self.result.setText(self.result.text()+'÷')
        self.result.adjustSize()
    def wfloat(self):
        ls.append('.')
        self.result.setText(self.result.text()+'.')
        self.result.adjustSize()

    def clear(self):
        ls.pop(len(ls)-1)
        a=str()
        for item in range(len(ls)):
            if ls[item]=='*':
                a+='×'
            elif ls[item]=='/':
                a+='÷'
            elif ls[item]=='.':
                a+='.'
            else:
                a+=ls[item]
        self.result.setText('')
        self.result.setText(self.result.text()+a)
        self.result.adjustSize()

    def clearall(self):
        ls.clear()
        self.lst.clear()
        self.number=''
        self.x=''
        self.answr=''
        self.result.setText('')
    
    def answer(self):
        self.lst=list()
        self.result.setText('')
        self.number=''
        for item in range(len(ls)):
            if ls[item] in '1234567890.':
                self.number+=ls[item]
            else:
                self.lst.append(self.number)
                self.lst.append(ls[item])
                self.number=''
        self.lst.append(self.number)
        q=1
        while q==1:
            if '*' not in self.lst and '/' not in self.lst:
                q=0
            for item in range(1,len(self.lst)-1):
                if self.lst[item]=='/':
                    self.lst[item-1]=float(self.lst[item-1])/float(self.lst[item+1])
                    self.lst[item]=float(self.lst[item-1])/float(self.lst[item+1])
                    self.lst[item+1]=float(self.lst[item-1])/float(self.lst[item+1])
                    self.lst.pop(item)
                    self.lst.pop(item)
                    break
                elif self.lst[item]=='*':
                    self.lst[item-1]=float(self.lst[item-1])*float(self.lst[item+1])
                    self.lst[item]=float(self.lst[item-1])*float(self.lst[item+1])
                    self.lst[item+1]=float(self.lst[item-1])*float(self.lst[item+1])
                    self.lst.pop(item)
                    self.lst.pop(item)
                    break
        k=1
        while k==1:
            if '+' not in self.lst and '-' not in self.lst:
                k=0
            for item in range(1,len(self.lst)-1):
                if self.lst[item]=='+':
                    self.lst[item-1]=float(self.lst[item-1])+float(self.lst[item+1])
                    self.lst[item]=float(self.lst[item-1])+float(self.lst[item+1])
                    self.lst[item+1]=float(self.lst[item-1])+float(self.lst[item+1])
                    self.lst.pop(item)
                    self.lst.pop(item)
                    break
                elif self.lst[item]=='-':
                    self.lst[item-1]=float(self.lst[item-1])-float(self.lst[item+1])
                    self.lst[item]=float(self.lst[item-1])-float(self.lst[item+1])
                    self.lst[item+1]=float(self.lst[item-1])-float(self.lst[item+1])
                    self.lst.pop(item)
                    self.lst.pop(item)
                    break
        self.answr=str(self.lst[0])
        t=0
        if self.answr[-2]=='.' and self.answr[-1]=='0':
            self.x=str(int(self.lst[0]))
        else:
            self.x=str(self.lst[0])
        self.result.setText(self.result.text()+self.x)
        self.result.adjustSize()

app=QApplication(sys.argv)
obj=Window()
obj.show()
sys.exit(app.exec_())
