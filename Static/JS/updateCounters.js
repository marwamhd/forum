function UpdatesLikesCounter(likes, dislikes) {
    var CounterDiv = document.getElementById("counterForLikes")
    CounterDiv.innerHTML = "Likes count: " + likes + " Dislikes count: " + dislikes
}

function UpdatesCommentLikesCounter(commentid, likes, dislikes) {
    var CounterDiv = document.getElementById("counterForLikes" + commentid)
    CounterDiv.innerHTML = "Likes count: " + likes + " Dislikes count: " + dislikes
}