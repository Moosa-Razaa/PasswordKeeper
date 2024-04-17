import { useState } from "react";
import login_image from "../../assets/left side image1.jpg"
import style from './login.module.css'
import LoginModal from "../../interfaces/login";
import axios from "axios";
// import 'antd/dist/antd.css';
// import { Button } from 'antd';

function Login() {
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [alertBorderColor, setalertBorderColor] = useState<string>("");

    function Email_OnChangeHandler(value: React.ChangeEvent<HTMLInputElement>) {
        setEmail(value.target.value);
    }

    function Pass_OnChange(pass: React.ChangeEvent<HTMLInputElement>) {
        setPassword(pass.target.value);
    }

    function handleKeyDown(event: React.KeyboardEvent<HTMLInputElement>) {
        if (event.key==="Enter") {
            LoginButton_OnClickHandler();
        }
    }

    function LoginButton_OnClickHandler() {
        const body: LoginModal = {
            Name: email,
            Password: password
        };

        axios.post("http://192.168.0.33:3000/login", body)
        .then((response) => {
            console.log("Response from server : ", response);
        })
        .catch((error) => {
            console.log("Error from API : ", error);
        });
    }

    const backgroundImageStyle = {
        backgroundImage: `url(${login_image})`
    };

    return (
        <div className={style["flex-container"]}>
            <div className={style["left"]} style={backgroundImageStyle}>
                <div className={style["heading"]}>
                    <h1> The Password <br /> Keeper</h1>
                </div>
            </div>
            <div className={style["content"]}>
                <div className={style["right"]}>
                    <h1 className="main-heading">Login</h1>
                    <div className={style["seperator"]}></div>
                    <p>Email: <input type="text" value={email} placeholder="Enter Email"
                        onChange={Email_OnChangeHandler}
                        style={{ borderColor: alertBorderColor }} id="border-element"
                        className={`${style["shake-animation"]} ${alertBorderColor === "red" ? style["shake-animation"] : " "}`} /></p>
                    <p>Password: <input type="password" value={password} placeholder="Enter Password" onChange={Pass_OnChange} onKeyDown={handleKeyPress} /></p>
                    <button onClick={LoginButton_OnClickHandler} className={style["login-button"]}>Login</button>
                </div>
            </div>
        </div>
    )
}

export default Login;