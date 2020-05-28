int sensors[]{A0,A1,A2,A3,A4,A5};
String names[]{"A0","A1","A2","A3","A4","A5"};
#define echoPin 11
#define trigPin 12

void setup()
{
  pinMode(trigPin, OUTPUT);
  pinMode(echoPin, INPUT);    
  Serial.begin(9600);
}

void loop()
{
  analogSensors();
  distancePing();
  delay(1000);
}

void analogSensors(){
  int numberOfSensors = sizeof(sensors)/sizeof(int);
  
  for (int i = 0; i < numberOfSensors; i++) {
    int value = analogRead(sensors[i]);
    Serial.print(names[i]);
    Serial.print(":");
    char valueString[5];
    sprintf(valueString,"%05d",value);
    Serial.print(valueString);
    Serial.println();
  }
}

void distancePing(){
  digitalWrite(trigPin, LOW); 
  delayMicroseconds(2); 
  digitalWrite(trigPin, HIGH);
  delayMicroseconds(10); 
  digitalWrite(trigPin, LOW);
  long duration = pulseIn(echoPin, HIGH);
  //Calculate the distance (in cm) based on the speed of sound.
  float distance = ((float)duration)/58.2;
  Serial.print("US:");
  Serial.println(distance);
}
