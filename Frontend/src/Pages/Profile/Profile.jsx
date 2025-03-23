import React, { use, useEffect, useState } from "react";
import "./Profile.css";
import Navbar from "../../Components/Navbar/Navbar";
import { getProfile } from "../../API/API";
import { Link } from "react-router-dom";
import { MdOutlineKeyboardBackspace } from "react-icons/md";

function Profile() {
    const [user, setUser] = useState({
        name: "",
        email: "",
        role: "",
        contact_number: "",
        libid: 0,
    });

    useEffect(() => {
        async function fetchData() {
            try {
                const res = await getProfile();
                const userData = res.data.data;
                console.log(userData);
                setUser({
                    name: userData.name,
                    email: userData.email,
                    role: userData.role,
                    contact_number: userData.contact_number,
                    libid: userData.lib_id,
                });
            } catch (error) {
                console.error("Token not found in localStorage");
            }
        }
        fetchData();
    }, []);

    return (
        <>
            <Navbar />
            <div className="profile-wrapper">
                <div className="top-wrapper">
                    <h1 className="heading">Profile</h1>
                    <div className="profile-container">
                        <div className="row">
                            <div className="box a">Name:</div>
                            <div className="box b">Email:</div>
                            <div className="box c">Contact Number:</div>
                            <div className="box d">Role:</div>
                            {user.libid !== 0 && (
                                <div className="box e">Lib ID</div>
                            )}
                        </div>
                        <div className="row">
                            <div className="box a-value">{user.name}</div>
                            <div className="box b-value">{user.email}</div>
                            <div className="box c-value">
                                {user.contact_number}
                            </div>
                            <div className="box d-value">{user.role}</div>
                            {user.libid !== 0 && (
                                <div className="box e-value">{user.libid}</div>
                            )}
                        </div>
                    </div>
                </div>
                <div className="child-back-btn">
                    <Link to="/home">
                        <MdOutlineKeyboardBackspace size={45} />
                    </Link>
                </div>
            </div>
        </>
    );
}

export default Profile;
