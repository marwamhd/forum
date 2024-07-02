// Function to retrieve post object by id
const getPostById = (postId) => {
    for (let i = 0; i < initialPosts.length; i++) {
        if (initialPosts[i].ID === postId) {
            return initialPosts[i];
        }
    }
    return null; // Return null if no post found with the given id
};

// Function to retrieve post object by id
const getPostByIdAndArray = (array, postId) => {

    for (let i = 0; i < array.length; i++) {
        if (array[i].ID === postId) {
            return array[i];
        }
    }
    return null; // Return null if no post found with the given id
};

function AddToThePost(postId, comment) {

    var contentDiv = document.querySelector('.Post');

    // Create a new comment element
    var newComment = document.createElement('div');
    newComment.classList.add('comment');
    newComment.innerHTML = `
        <div class="user">Comment by You</div>
        <div class="content">
            ${comment}
        </div>`;

    var formElement = document.getElementById('commentForm-' + postId);

    contentDiv.insertBefore(newComment, formElement.parentNode);

    //add to initialPosts
    initialPosts.forEach(function(post) {
        if (post.ID === postId) {
            post.comments.push({
                U_ID: 'You',
                comment: comment
            });
        }
    });

}