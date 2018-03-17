var markdowner = function(){
    var content = document.getElementById('markdown-content-in').value;
    var title = document.getElementById('markdown-title-in').value;
    var markdownOut = document.getElementById('markdown-content-out');
    var markdownTitleOut = document.getElementById('markdown-title-out');

    markdownTitleOut.innerHTML = '<h1>'+title+'</h1>';
    markdownOut.innerHTML = marked(content);      
}

var url = '/api/markdown'
fetch(
    url ,
    {method: 'POST'}
).then(
    res => res.json()
).then(data => {
    console.log("test")
    ReactDOM.render(
        MDMaster(data),
        document.getElementById('markdown')
    )
}).catch( e => {
    ReactDOM.render(
        <div id="markdown-out">
            <div id="markdown-title-out"><h1>タイトル</h1></div>
            <div id="markdown-content-out"><h1>各種設定項目</h1></div>
        </div>,
        document.getElementById('markdown')
    )
});

var MDMaster = (props) => {
    data = props.items
    return <div id="markdown-out">
            <MDTitle items={data.Title} />
            <MDContent items={data.Content} />
        </div>
}

var MDTitle = (props) => {
    var item = props.items
    return <div id="markdown-title-out"><h1>{item}</h1></div>
    
}


// MDContent (一応引数は、contentとmdtitleにしたい)
var MDContent = (props) => {
    var item = props.items
    return <div id="markdown-content-out"><h1>{item}</h1></div>
}

const DefaultData = {
    Title: "タイトル要素",
    Content: "各種設定"
}


