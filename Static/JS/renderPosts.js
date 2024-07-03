function renderPosts(array) {
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    // Start building the HTML content
    var htmlContent = `<div class="contentpost"><div class="posts">`;

    // Iterate through initialPosts and construct each post HTML
    array.forEach(function(post) {
        htmlContent += `
            <div class="post">
                <div class="user">Post created by ${post.Username}</div>
                <div class="title">
                    <h2>${post.Title}</h2>
                </div>
                <div class="content">
                    ${post.Post}
                </div>
                <div>
                <a href="#" class="viewbtn" onclick="viewPost( ${post.ID}, initialPosts)"> view </a>
                </div>
            </div> `;
    });

    // Complete the HTML content with closing div tags
    htmlContent += `</div></div>`;

    // Set the entire constructed HTML content to contentDiv.innerHTML
    contentDiv.innerHTML = htmlContent;

    // Optionally, you can log the constructed HTML to verify
}

function renderPostsByID(uid) {
    var contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear existing content

    // Start building the HTML content
    var htmlContent = `<div class="contentpost"><div class="posts">`;

    // Iterate through initialPosts and construct each post HTML
    initialPosts.forEach(function(post) {
        // Check if the post's u_id matches the given uid parameter
        if (post.U_ID === uid) {
            htmlContent += `
                <div class="post">
                    <div class="user">Post created by ${post.Username}</div>
                    <div class="title">
                        <h2>${post.Title}</h2>
                    </div>
                    <div class="content">
                        ${post.Post}
                    </div>
                    <div>
                        <a class="viewbtn" href="#" onclick="viewPost(${post.ID}, initialPosts)">View</a>
                    </div>
                </div> `;
        }
    });

    // Complete the HTML content with closing div tags
    htmlContent += `</div></div>`;

    // Set the entire constructed HTML content to contentDiv.innerHTML
    contentDiv.innerHTML = htmlContent;

}