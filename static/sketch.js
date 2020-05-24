var rects1;
var points;
var queryPoints;

var x;
var y;
var colors =[];
var greys = {}
var runQuery = false;

function preload() {
  colors = [color(141,211,199),
    color(255,255,179),
    color(190,186,218),
    color(251,128,114),
    color(128,177,211),
    color(253,180,98),
    color(179,222,105),
    color(252,205,229),
    color(217,217,217),
    color(188,128,189),
    color(204,235,197),
    color(255,237,111)]

    greys = {
      '0.09765625' : color(255),
      '0.1953125' : color(0,0, 0, 25),
      '0.390625' : color(0,0, 0, 50),
      '0.78125' : color(0,0, 0, 75),
      '1.5625' : color(0,0, 0, 100),
      '3.125' : color(0,0, 0, 126),
      '6.25' : color(0,0, 0, 150),
      '12.5' : color(0,0, 0, 175),
      '25' : color(0,0, 0, 200),
      '50' : color(0,0, 0, 225),
      '100' : color(0,0, 0, 255),
    }

    loadJSON("/rects", callbackRects);
    loadJSON("/query?x=-4.415&y=54.65&w=13&h=11", callbackPoints);
}

function callbackRects(data) {
    rects1 = data;
}

function callbackPoints(data) {
    points = data;
}

function callbackQuery(data) {
    queryPoints = data;
}

function setup() {
  
  background(51)  
  createCanvas(1000,850)
  extraCanvas = createCanvas(1000,850);
  extraCanvas
}

function draw() {

 	background(255);
  rectMode(CENTER);
  stroke(0)
  translate(10,10)
  fill(255)

  var widths = []
  
  
  for(var i = 0; i < rects1.length; i++) {
    var r = rects1[i]
   
    x1 = map(r.X, -6, 2, 0, 400)
    y1 = map(r.Y, 50, 59, 800, 0)
    w1 = map(r.W, 0, 8, 0, 400)
    h1 = map(r.H, 0, 9, 0, 800)
    //fill(greys[w1]);
    if (!widths.includes(w1)) {
      widths.push(w1)
    }
   
    //noStroke()
    rect(x1, y1, w1, h1)
  
    //widths = widths.sort(function(a, b){return a - b});
  }
  
  fill(255)
  for(var i = 0; i < points.length; i++) {
    var r = points[i] 
    r.L < colors.length ? fill(colors[r.L]) : fill(255)
    x = map(r.X, -6, 2, 0, 400)
    y = map(r.Y, 50, 59, 800, 0)
    //ellipse(x,y, 5, 5)
    stroke(r.L)
    point(x,y)
    
  }

  if(queryPoints) {
    fill(0,0, 255,125)
    for(var i = 0; i < queryPoints.length; i++) {
      var r = queryPoints[i]
      x = map(r.X, -6, 2, 0, 400)
      y = map(r.Y, 50, 59, 800, 0)
      ellipse(x, y, 10, 10)
    }
  }
  push()
  noStroke()
  fill(0)
  textSize(32);
  if(queryPoints) {
    text("Number of points found : " + queryPoints.length, 410, 40);
  }
  pop()
  
  fill(0,255,0,100);
  
  rect(mouseX, mouseY, 50, 50);   
}

function mousePressed() {
  runQuery = true;
}

function mouseReleased() {
  runQuery = false;
}

function mouseDragged() {
    x = map(mouseX,0, 400, -6, 2)
    y = map(mouseY,800, 0, 50, 59)
    w = map(50,0, 400, 0,8)
    h = map(50,0, 400, 0,8)

    w = map(50, 0, 400, 0, 8)
    h = map(50, 0, 800, 0, 9)

    if(runQuery) {
      loadJSON("/query?x=" + x + "&y=" + y + "&w=" + w + "&h=" + h, callbackQuery);
    }
}