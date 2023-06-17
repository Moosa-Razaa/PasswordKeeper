import { useState } from "react";
import login_image from "../../assets/20945597.jpg"
import logo from "../../assets/images.png"
import style from './login.module.css'
import alert_image from '../../assets/alert_image.svg'


function Login()
{
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [alertImageHidden, setHidden] = useState<boolean>(true);

    function Email_OnChangeHandler(value : React.ChangeEvent<HTMLInputElement>)
    {
        setEmail(value.target.value);
    }
    function Pass_OnChange(pass : React.ChangeEvent<HTMLInputElement>) {
        setPassword(pass.target.value);   
    }
    
    function ValidateEmail(emailValue: string) : boolean
    {
        const emailRegex: RegExp = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(emailValue);
    }
    
    function LoginButton_OnClickHandler() {
        const returnedValue:boolean = ValidateEmail(email);
        setHidden(returnedValue);
    }

    return(
        <div className={style["main"]}>
             <div className={style["navbar"]}>
            <img src={logo} alt="" className={style["logo"]} />
                <img src={login_image} className={style["login_image"]} alt="" />
            </div>
            <div className={style["login_form"]}>
                <div className={style["content"]}>
                    <h1 className={style["heading"]}><span>The</span> Password <br /> Keeper</h1>
                    <h1 className={style["login"]}>LOGIN</h1>
                    <div className={style["seperator"]}></div>
                    <p className={style["ptag1"]}>Email :</p>
                    <input value={email} onChange={Email_OnChangeHandler} type="email" name="" className={style["user"]}  placeholder="Enter Email" />
                    <img  hidden={alertImageHidden} src={alert_image} className={style["email_alert_image"]} alt="" />
                    <p className={style["ptag2"]}>Password :</p>
                    <input value={password} type="password" className={style["password"]} onChange={Pass_OnChange} placeholder="Enter Password" />
                    <button onClick={LoginButton_OnClickHandler}>Login</button>
                    <p className={style["bottom"]}>don't have an account? <a href="">signup</a></p>
                </div>
            </div>
        </div>

    )
}

export default Login;