function submitLike() {
    const form = document.getElementById('likeForm');
    const formData = new FormData(form);


    if (useridentification==0){
        alert("not authorized")
        return false;
    }


    fetch('/addlike', {
            method: 'POST',
            body: formData
        })
        .then(response => response.json())
        .then(data => {
            console.log('Server response:', data);
            // Handle the response accordingly
            if (data.success) {
                console.log('Like/dislike submitted successfully');
                UpdatesLikesCounter(data.likes, data.dislikes)
                    // Optionally, update UI or perform additional actions
            } else {
                console.error('Error submitting like/dislike:', data.error);
            }
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
}

function submitComment(postId) {
    const form = document.getElementById(`commentForm-${postId}`);
    const formData = new FormData(form);


    if (useridentification==0){
        alert("not authorized")
        return false;
    }

    fetch('/addcomment', {
            method: 'POST',
            body: formData
        })
        .then(response => response.text())
        .then(text => {
            console.log('Server response:', text);
            try {
                const data = JSON.parse(text);
                if (data.success) {
                    console.log('Comment added:', data);
                    // Reload the page to view the newly added comment

                    // AddToThePost(postId, formData.get('comment'));  

                    viewPost(postId, data.posts)

                } else {
                    console.error('Error adding comment:', data.error);
                }
            } catch (error) {
                console.error('JSON parsing error:', error);
            }
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
}

function submitCommentLike(commentid) {
    const form = document.getElementById('CommentlikeForm' + commentid);
    const formData = new FormData(form);


    if (useridentification==0){
        alert("not authorized")
        return false;
    }


    fetch('/addCommentlike', {
            method: 'POST',
            body: formData
        })
        .then(response => response.json())
        .then(data => {
            console.log('Server response:', data);
            // Handle the response accordingly
            if (data.success) {
                console.log('Like/dislike submitted successfully');
                UpdatesCommentLikesCounter(commentid, data.likes, data.dislikes)
                    // Optionally, update UI or perform additional actions
            } else {
                console.error('Error submitting like/dislike:', data.error);
            }
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });

}


function GetLikedPosts() {
    fetch('/likedpost', {
            method: 'POST',
            body: ""
        })
        .then(response => response.text())
        .then(text => {
            console.log('Server response:', text);
            try {
                const data = JSON.parse(text);
                if (data.success) {
                    console.log('Comment added:', data);
                    // Reload the page to view the newly added comment
                    renderPosts(data.posts);

                } else {
                    console.error('Error adding comment:', data.error);
                }
            } catch (error) {
                console.error('JSON parsing error:', error);
            }
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
}
