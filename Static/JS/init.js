window.onload = function() {
    let pageIn = localStorage.getItem("pageIn");
    if (pageIn == null) {
        pageIn = 'home'
    }

    if (pageIn.includes('viewPost')) {
        viewPost(Number(pageIn.split(' ')[1]), initialPosts)
    } else {
        changeContent(pageIn)
    }

    const fragment = window.location.hash.substring(1); 

    if (fragment.length != 0) {
        window.location.href = '/'
    }

    // alert(pageIn)
};