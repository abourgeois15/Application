import { useEffect, useState } from "react";
import { useErrorHandler } from "react-error-boundary";

export const useApi = (callBack, initial, ...args) => {
  const [state, setState] = useState(initial);
  const handleError = useErrorHandler();
  console.log(args)
  const fetchData = async () => {
    try {
      const res = await callBack(...args);
      setState(res);
    } catch (error) {
        handleError(error);
    }
  };

  useEffect(() => {
    fetchData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [JSON.stringify(args)]
  );

  return {state, setState};
};