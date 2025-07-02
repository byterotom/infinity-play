const voteKey = `infinity_vote_${gameId}`;

function disableButtons() {
    document.getElementById('like-btn').disabled = true;
    document.getElementById('dislike-btn').disabled = true;
}

function handleVote(voteType) {
    localStorage.setItem(voteKey, voteType);
    disableButtons();
}

// Disable buttons on page load if already voted
document.addEventListener('DOMContentLoaded', function () {
    if (localStorage.getItem(voteKey)) {
        disableButtons();
    }
});


document.getElementById('fullscreenBtn').addEventListener('click', function () {
    const gameContainer = document.getElementById('gameContainer');
    const gameObject = document.getElementById('gameObject');

    if (!document.fullscreenElement) {
        // Enter fullscreen
        if (gameContainer.requestFullscreen) {
            gameContainer.requestFullscreen();
        } else if (gameContainer.webkitRequestFullscreen) {
            gameContainer.webkitRequestFullscreen();
        } else if (gameContainer.mozRequestFullScreen) {
            gameContainer.mozRequestFullScreen();
        } else if (gameContainer.msRequestFullscreen) {
            gameContainer.msRequestFullscreen();
        }

        // Scale game object to fit fullscreen
        gameObject.style.width = '100vw';
        gameObject.style.height = '100vh';
    } else {
        // Exit fullscreen
        if (document.exitFullscreen) {
            document.exitFullscreen();
        } else if (document.webkitExitFullscreen) {
            document.webkitExitFullscreen();
        } else if (document.mozCancelFullScreen) {
            document.mozCancelFullScreen();
        } else if (document.msExitFullscreen) {
            document.msExitFullscreen();
        }

        // Reset game object size
        gameObject.style.width = '800px';
        gameObject.style.height = '500px';
    }
});

// Handle fullscreen change events
document.addEventListener('fullscreenchange', handleFullscreenChange);
document.addEventListener('webkitfullscreenchange', handleFullscreenChange);
document.addEventListener('mozfullscreenchange', handleFullscreenChange);
document.addEventListener('msfullscreenchange', handleFullscreenChange);

function handleFullscreenChange() {
    const gameObject = document.getElementById('gameObject');
    if (!document.fullscreenElement) {
        // Reset size when exiting fullscreen
        gameObject.style.width = '800px';
        gameObject.style.height = '500px';
    }
}