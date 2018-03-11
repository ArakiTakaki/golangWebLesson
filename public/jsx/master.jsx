
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

ReactDOM.render(
    <h1>Hello, world!</h1>,
    document.getElementById('root')
  );
