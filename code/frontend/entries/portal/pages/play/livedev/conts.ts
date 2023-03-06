export const ServerJS = "server.js";
export const ClientJS = "client.js";

const client = `

liveapp.register("myapp.main", myApp);
liveapp.loadLib("chartjs.js");
liveapp.loadLib("mnop.js");

const myApp = async (opts) =>  {
  
  
  opts.utils.AutoForm({
    data: [1, 2, 3]
    on_submit: (data) => {
    	console.log("@data", data)
  	}
  })
  
  
  console.log(opts);
}
`;
