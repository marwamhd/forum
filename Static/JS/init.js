window.onload = function() {
    renderPosts(initialPosts);
    var footer = document.getElementById("cats")
    footer.innerHTML = `        
    <form action="/" method="get">
        <legend>Categories</legend>
        <label>
            <input type="checkbox" name="cat" value="cat1"> cat1
        </label><br>
        <label>
            <input type="checkbox" name="cat" value="cat2"> cat2
        </label><br>
        <label>
            <input type="checkbox" name="cat" value="cat3"> cat3
        </label><br>
        <button type="submit">Filter</button>
    </form>`
};