import React from "react";
import "./Card.css";

function Card({ bookImage = "Book Image", desc = "Description" }) {
	return (
		<>
			<div className="card">
				<div className="book-image">{bookImage}</div>
				<div className="book-desc">{desc}</div>
				<button className="btn-view-book">View Book</button>
				<button className="btn-request">request</button>
			</div>
		</>
	);
}

export default Card;
