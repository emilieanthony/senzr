import "./card.css";

function Card(props) {
  let classes = "card";
  if (props.className) {
    classes += " " + props.className;
  }
  return <div className={classes}>{props.children}</div>;
}

export default Card;
