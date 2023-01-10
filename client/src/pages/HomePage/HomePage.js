import "./homepage.css";

export const HomePage = () => {

  return (
    <div data-cy="home-page" className="homePageWrapper">
      <div className="contentContainer">
        <h1 data-cy="header" className="header">Welcome to T&S</h1>
        <p data-cy="motto" className="motto">Let's make it Possible!!!</p>
      </div>

      <div className="linkContainer">
        <a data-cy="link" className="link" href="/fullItems">
          Items
        </a>
      </div>
      <div className="linkContainer">
        <a data-cy="link" className="link" href="/fullMachines">
          Machines
        </a>
      </div>
    </div>
  );
};