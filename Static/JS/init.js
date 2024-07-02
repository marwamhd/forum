window.onload = function() {
    renderPosts(initialPosts);
    var footer = document.getElementById("cats")
    footer.innerHTML = `        
    <form action="/" method="get">
        <label>
            <input type="checkbox" name="cat" value="cat1"> cat1
        </label><hr>
        <label>
            <input type="checkbox" name="cat" value="cat2"> cat2
        </label><hr>
        <label>
            <input type="checkbox" name="cat" value="cat3"> cat3
        </label><hr>
        <button type="submit">Filter</button>
    </form>`
};