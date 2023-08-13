import "./App.css";
import { useCookies } from "react-cookie";

const logout = `${process.env.REACT_APP_BACKEND}/api/auth/google/logout`;
const login = `${process.env.REACT_APP_BACKEND}/api/auth/google/login`;

function App() {
  const [cookies, removeCookies] = useCookies(["token"]);

  const onClick = () => {
    window.location.href = logout;
    removeCookies("token");
  };

  if (cookies?.token !== "undefined") {
    return (
      <div className="App">
        <form
          onSubmit={(e) => {
            let objective = e.target.objective.value;
            let translate = e.target.translate.value;

            console.log({ objective, translate });

            fetch(`${process.env.REACT_APP_BACKEND}/api/v1/okr-generator`, {
              method: "POST",
              mode: "cors",
              headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${cookies?.token}`,
              },
              body: JSON.stringify({
                objective,
                translate,
              }),
            })
              .then((res) => res.json())
              .then((res) => console.log(res));

            e.preventDefault();
          }}
        >
          <input type="text" name="objective" placeholder="objective" />
          <input type="text" name="translate" placeholder="translate" />
          <input type="submit" />
        </form>
        <button onClick={onClick}>logout</button>
      </div>
    );
  } else {
    return (
      <div>
        <a href={login}>Login</a>
      </div>
    );
  }
}

export default App;
