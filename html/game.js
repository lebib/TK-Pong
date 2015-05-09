// WORLD

var stage = document.getElementById("gameCanvas");
stage.width = STAGE_WIDTH;
stage.height = STAGE_HEIGHT;

var context = stage.getContext("2d");
context.fillStyle = "grey";
context.font = GAME_FONTS;

///////////////////////////////////////////////////////////////////////////////

// PRELOAD
var paddle_p1 = new Image(),
    paddle_p2 = new Image(),
    ball = new Image();

paddle_p1.ready = false;
paddle_p2.ready = false;
ball.ready = false;

paddle_p1.onload = setAssetReady;
paddle_p2.onload = setAssetReady;
ball.onload = setAssetReady;

paddle_p1.src = PADDLE_P1_IMAGE;
paddle_p2.src = PADDLE_P2_IMAGE;
ball.src = BALL_IMAGE;

function setAssetReady()
{
    this.ready = true;
}

//

context.fillRect(0, 0, stage.width, stage.height);
context.fillStyle = "#000";
context.fillText("LOADING", 10, 10);
var preloader = setInterval(preloading, TIME_PER_FRAME);
var gameloop;

function preloading()
{
    if(paddle_p1.ready)
    {
        clearInterval(preloader);
        gameloop = setInterval(update, TIME_PER_FRAME);
    }
}

var paddle_p1_x = 0,
    paddle_p1_y = PADDLE_P1_Y,
    paddle_p2_x = 0,
    paddle_p2_y = PADDLE_P2_Y
    ball_x = 50,
    ball_y = 50;

function update()
{
    context.fillStyle = "grey";
    context.fillRect(0, 0, stage.width, stage.height);
    context.drawImage(paddle_p1, paddle_p1_x, paddle_p1_y)
    context.drawImage(paddle_p2, paddle_p2_x, paddle_p2_y)
    context.drawImage(ball, ball_x, ball_y);
}

///////////////////////////////////////////////////////////////////////////////
