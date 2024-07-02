window.onload = function() {
    renderPosts(initialPosts);
    var footer = document.getElementById("cats")
    footer.innerHTML = `        
    <form action="/" method="get">
        <label>
            <input type="checkbox" name="cat" value="cat1"> Annoucements
        </label><hr>
        <label>
            <input type="checkbox" name="cat" value="cat2"> Events
        </label><hr>
        <label>
            <input type="checkbox" name="cat" value="cat3"> Questions
        </label><hr>
        <button type="submit">Filter</button>
    </form>`
};