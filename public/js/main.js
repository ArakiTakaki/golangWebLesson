


var NavItems = props => {
    const items = props.items

    // アローを使ってfor文的なの書く
    const listItems = items.map((item) =>
        <li key={item.Key.toString()}>
            <a href={item.URL}>{item.Name}</a>
        </li>
    );
    // 先にループ処理をさせておいてから使う
    return <ul>
        {listItems}
    </ul>
}

fetch(`/api/meta`,{method: 'get'}).then(res => res.json()).then(data => {
    ReactDOM.render(
        <div id="left-nav">
            <input type="checkbox" id="nav-button" />
            <label htmlFor="nav-button">MENU</label>
            <section>
                    <NavItems items={data.Page} />        
            </section>
        </div>,
        document.getElementById('nav')
    );
});

// ItemDetailのコンポーネント定義
// const ItemDetail = props => {
//     const item = props.item;
//     return <li className={'item' + item.stock === 0 ? ' soldout' : ''}>
//         <div className="item-name">{item.name}</div>
//         <div className="item-price">{item.price}</div>
//     </li>;
// };

// fetch('/js/test.json').then(res => res.json()).then(data => {
//     ReactDOM.render(
//         <ItemList items={data.items} />, // これを
//         document.getElementById('container') // ここにレンダリングしろ
//     );
// });