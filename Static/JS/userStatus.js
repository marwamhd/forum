function GetIfUserLiked(pid) {
    const requestData = {
        pid: pid
    };

    fetch('/diduserlike', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        })
        .then(response => response.json())
        .then(data => {
            console.log('Server response:', data);
            // Handle the response accordingly
            if (data.success) {
                console.log('User liked status retrieved successfully');
                switch (data.userl) {
                    case 0:
                        document.getElementById('dislike').checked = true; // Assuming 'disliked' is the ID of the dislike radio button
                        break;
                    case 1:
                        document.getElementById('like').checked = true; // Assuming 'liked' is the ID of the like radio button
                        break;
                    default:
                        console.error('Unknown user liked status:', data.userl);
                        break;
                }
            } else {
                console.error('Error retrieving user liked status:', data.error);
            }
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
}

function GetIfUserLikedComment(pid, cid) {
    const requestData = {
        pid: pid,
        cid: cid
    };

    fetch('/diduserlikecomment', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        })
        .then(response => response.json())
        .then(data => {
            console.log('Server response:', data);
            // Handle the response accordingly
            if (data.success) {
                console.log('User liked status retrieved successfully');
                switch (data.userl) {
                    case 0:
                        document.getElementById('dislike' + cid).checked = true; // Assuming 'disliked' is the ID of the dislike radio button
                        break;
                    case 1:
                        document.getElementById('like' + cid).checked = true; // Assuming 'liked' is the ID of the like radio button
                        break;
                    default:
                        console.error('Unknown user liked status:', data.userl);
                        break;
                }
            } else {
                console.error('Error retrieving user liked status:', data.error);
            }
        })
        .catch(error => {
            console.error('Fetch error:', error);
        });
}