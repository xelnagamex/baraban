#include <LiquidCrystalRus.h>

// RS, E, DB4, DB5, DB6, DB7
LiquidCrystalRus lcd(12, 11, 5, 4, 3, 2);

unsigned char data = 0;

void setup()
{
  int i;
    lcd.begin(16, 2);
    lcd.setCursor(4, 0);
    lcd.print("Loading");
    lcd.setCursor(0, 1);
    for (i=0;i<15;i++) {
      lcd.print("#");
      delay(40);
    }
    Serial.begin(9600);
    lcd.clear();
}

void loop() {

  if (Serial.available()) {
    delay(100);
    lcd.clear();
    while (Serial.available() > 0) {
      data = Serial.read();
      //if (data == '\x0A') { break;}
      if (data == '\x01') {
        lcd.setCursor(0, 1);
        continue;
      }
      lcd.write(data);
    }
  }
}
