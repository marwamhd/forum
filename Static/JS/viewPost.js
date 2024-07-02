function viewPost(pid, array) {
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    GetIfUserLiked(pid)

    let a = getPostByIdAndArray(array, pid);

    // Start building the HTML content for the single post view
    var htmlContent = `
        <div class="contentpost">
            <div class="posts">
                <div class="post">
                    <div class="user">Post created by ${a.Username}</div>
                    <div class="title">
                        <h2>${a.Title}</h2>
                    </div>
                    <div class="content">
                        ${a.Post}
                    </div>
                    <form id="likeForm" action="/addlike">
                        <input onclick="submitLike()" type="radio" id="like" name="like" value="1">
                        <label for="like">Like</label><br>
                        <input onclick="submitLike()" type="radio" id="dislike" name="like" value="0">
                        <label for="dislike">Dislike</label><br>
                        <input onclick="submitLike()" type="radio" id="prefernottosay" name="like" value="2">
                        <label for="prefernottosay">Prefernottosay</label>
                         <input type="hidden" name="pid" value="${a.ID}">
                    </form>

                    <div id="counterForLikes" > Likes count: ${a.Likes} Dislikes count: ${a.Dislikes} </div>

                `;

    // Iterate through comments and add HTML for each comment
    a.Comments.forEach(function(comment) {
        GetIfUserLikedComment(a.ID, comment.ID)
        console.log(comment)
        htmlContent += `
            <div class="comment">
                <div class="user">Comment by ${comment.U_ID}</div>
                <div class="content">
                    ${comment.Comment}
                </div>
                    <form id="CommentlikeForm${comment.ID}" action="/addCommentlike">
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="like${comment.ID}" name="like${comment.ID}" value="1">
                        <label for="like${comment.ID}">Like</label><br>
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="dislike${comment.ID}" name="like${comment.ID}" value="0">
                        <label for="dislike${comment.ID}">Dislike</label><br>
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="prefernottosay${comment.ID}" name="like${comment.ID}" value="2">
                        <label for="prefernottosay${comment.ID}">Prefernottosay</label>
                        <input type="hidden" name="cid" value="${comment.ID}">
                        <input type="hidden" name="pid" value="${a.ID}">
                    </form>

                    <div id="counterForLikes${comment.ID}" > Likes count: ${comment.Likes} Dislikes count: ${comment.Dislikes} </div>
            </div>`;
    });

    // Add HTML for the comment form
    htmlContent += `
            <div>
       <form id="commentForm-${a.ID}">
            <div>
                <label for="comment">Comment</label><br>
                <textarea name="comment" id="comment"></textarea><br>
                <input type="hidden" name="pid" value="${a.ID}">
                <button type="button" onclick="submitComment(${a.ID})">Add comment</button>
            </div>
        </form>
            </div>
        </div>
    </div>`;

    // Set the constructed HTML content to contentDiv.innerHTML
    contentDiv.innerHTML = htmlContent;

    // Log the post data for debugging purposes
}