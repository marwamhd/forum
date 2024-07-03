function viewPost(pid, array) {
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    SetPageRemember('viewPost ' + pid)

    GetIfUserLiked(pid)

    let a = getPostByIdAndArray(array, pid);
    let str = ""
    let str1 = ""
    let str2 = ""
    let str3 = ""

    if (a.Cat1 == 1) {
        str1 = "Annoucements"
    }
    if (a.Cat2 == 1) {
        str2 = "Events"
    }
    if (a.Cat3 == 1) {
        str3 = "Questions"
    }
    if (str1 != "" && str2 != "" && str3 != "") {
        str = str1 + ", " + str2 + " and " +str3
    } else if (str1 != "" && str2 != "") {
        str = str1 + " and " + str2
    } else if (str1 != "" && str3 != "") {
        str = str1 + " and " + str3
    } else if (str2 != "" && str3 != "") {
        str = str2 + " and " + str3
    } else if (str1 != ""){
        str = str1
    }else if (str2 != ""){
        str = str2
    }else if (str3 != ""){
        str = str3
    }


    // Start building the HTML content for the single post view
    var htmlContent = `
        <div class="contentpost">
            <div class="posts">
                <div class="post">
                    <span class="user">Post created by ${a.Username}</span>
                    <span class="catStr">Category: ${str}</span>
                    <div class="title">
                        <h2>${a.Title}</h2>
                    </div>
                    <div class="content">
                        ${a.Post}
                    </div>
                    <form id="likeForm" action="/addlike">
                        <input onclick="submitLike()" type="radio" id="like" name="like" value="1" ${useridentification == 0 ? 'disabled' : ''}>
                        <label for="like">Like</label>
                        <input onclick="submitLike()" type="radio" id="dislike" name="like" value="0" ${useridentification == 0 ? 'disabled' : ''}>
                        <label for="dislike">Dislike</label>
                        <input onclick="submitLike()" type="radio" id="prefernottosay" name="like" value="2" ${useridentification == 0 ? 'disabled' : ''}>
                        <label for="prefernottosay">Prefernottosay</label>
                         <input type="hidden" name="pid" value="${a.ID}">
                    </form>

                    <div class="endofpost" id="counterForLikes" > Likes count: ${a.Likes} Dislikes count: ${a.Dislikes} </div>
                    <hr>

                `;

    // Iterate through comments and add HTML for each comment
    a.Comments.forEach(function(comment) {
        GetIfUserLikedComment(a.ID, comment.ID)
        htmlContent += `
            <div class="comment">
                <div class="user">Comment by ${comment.Username}</div>
                <div class="content">
                    ${comment.Comment}
                </div>
                    <form id="CommentlikeForm${comment.ID}" action="/addCommentlike">
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="like${comment.ID}" name="like${comment.ID}" value="1"${useridentification == 0 ? 'disabled' : ''}>
                        <label for="like${comment.ID}">Like</label>
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="dislike${comment.ID}" name="like${comment.ID}" value="0"${useridentification == 0 ? 'disabled' : ''}>
                        <label for="dislike${comment.ID}">Dislike</label>
                        <input onclick="submitCommentLike(${comment.ID})" type="radio" id="prefernottosay${comment.ID}" name="like${comment.ID}" value="2"${useridentification == 0 ? 'disabled' : ''}>
                        <label for="prefernottosay${comment.ID}">Prefernottosay</label>
                        <input type="hidden" name="cid" value="${comment.ID}">
                        <input type="hidden" name="pid" value="${a.ID}">
                    </form>

                    <div  class="endofpost" id="counterForLikes${comment.ID}" > Likes count: ${comment.Likes} Dislikes count: ${comment.Dislikes} </div>
                    <hr>
            </div>`;
    });

    // Add HTML for the comment form
    htmlContent += `
            <div>
       <form id="commentForm-${a.ID}">
            <div>
                <label for="comment" class="commentlab">Comment</label><br>
                <textarea name="comment" id="comment" class="fixed-size-textarea" maxlength="500" ${useridentification == 0 ? 'disabled' : ''}></textarea><br>
                <input type="hidden" name="pid" value="${a.ID}">
                <button class="sameAsLog" type="button" onclick="submitComment(${a.ID})" ${useridentification == 0 ? 'disabled' : ''}>Add comment</button>
            </div>
        </form>
            </div>
        </div>
    </div>`;

    // Set the constructed HTML content to contentDiv.innerHTML
    contentDiv.innerHTML = htmlContent;


    var catsDiv = document.getElementById('cats');
    catsDiv.innerHTML = "";

    var fbuttom = document.getElementById("fbtn")
    fbuttom.innerHTML = ""


    // Log the post data for debugging purposes
}