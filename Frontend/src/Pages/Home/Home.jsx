import React from "react";
import "./Home.css";
import Navbar from "../../Components/Navbar/Navbar";
import Card from "../../Components/Card/Card";

export default function Home() {
	return (
		<>
			<Navbar />
			<div>Home</div>
			<div className="book-cards">
				<Card bookImage="Book Image 1" desc="Description 1" />
				<Card />
			</div>
		</>
	);
}
