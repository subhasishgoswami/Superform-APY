package server

import (
	"html/template"
)

func parseIndexTemplate() (*template.Template, error) {
	return template.New("index").Parse(
		`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Superform APY</title>
  <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
  <style>

    body {
      background-color: black;
      font-family: Arial, sans-serif;
      color: white;
    }

    .heading {
        font-size: 2.25rem;
        line-height: 2.5rem;
        font-weight: 600;
        color: white;
        display: flex;
        padding-top: 40px;
        align-items: center;
        justify-content: center;
    }
    
    .container {
      display: grid;
      grid-template-columns: repeat(2, 1fr);
      gap: 20px;
      margin: 20px;
    }

    .containerFull {
        margin: 20px;
    }
    
    .box {
      border-radius: 10px;
      padding: 20px;
      background-color: #2d2e35;
      box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5);
	  overflow:hidden;
	  overflow-y:scroll;
	  height: 90vh;
	  
    }
    
    .box h2 {
      font-size: 1.5rem;
      margin-top: 0;
      color: #00ed76;
    }
    
    .list {
      margin-top: 10px;
	  overflow:hidden;
	  overflow-y:scroll;
    }
    
    table {
      width: 100%;
      border-collapse: collapse;
      margin-top: 10px;
	  padding-bottom: 50px;
    }
    
    th, td {
      padding: 5px;
	  font-size: 1rem;
      text-align: left;
    }
    
    th {
      border-bottom: 1px solid white;
    }
    
    .chart {
      margin-top: 20px;
      height: 70vh;
      display: flex;
      justify-content: center;
    }

    .finePrint {
        padding-top: 10px;
        display: flex;
        justify-content: center;
        color: #2d2e35;
    }

    .footer {
        margin-top: 80px;
        padding-bottom: 30px;
        padding-top: 30px;
        background-color: #18191a;
        display: flex;
        justify-content: center;
    }

  </style>
</head>
<body>
<div class="heading">
    <img src="https://pon-app.netlify.app/static/media/icon-builder.7fb453f0.svg" alt="icon" style="padding: 10px;">
    Superform APY
</div>

  
<div class="container">
    
    <div class="box">
    <h2>API Status</h2>
    <div class="list">
    <table>
    <tbody>
        <tr>
        <td>User</td>
        <td>{{ .User }}</td>
        </tr>
        <tr>
        <td>Server Status</td>
        <td>OK</td>
        </tr>
        <tr>
        <td>Balance</td>
        <td>{{ .Balance }}</td>
        </tr>
        <tr>
        <td>Total Deposits</td>
        <td>{{ .TotalDeposits }}</td>
        </tr>
        <tr>
        <td>Total Withdrawals</td>
        <td>{{ .TotalWithdrawals }}</td>
        </tr>
        <tr>
        <td>Total Assets For Redemtion</td>
        <td>{{ .Redemption }}</td>
        </tr>
        <tr>
        <td>Total Earnings</td>
        <td>{{ .Earnings }}</td>
        </tr>
        <tr>
        <td>Total Profit</td>
        <td>{{ .Profit }}</td>
		</tr>
		<td>APR</td>
        <td>{{ .APR }} %</td>
        </tr>
    </tbody>
    </table>
    </div>
    </div>
</div>
<div class="heading">
    <img src="https://pon-app.netlify.app/static/media/icon-builder.7fb453f0.svg" alt="icon" style="padding: 10px;">
    Deposits And Withdrawals
</div>

<div class="container">    
	<div class="box">
      <h2>User Deposits</h2>
      <div class="list">
        <table>
          <thead>
            <tr>
			  <th>Block</th>
              <th>Asset Deposited</th>
              <th>Share Share Recieved</th>
              
            </tr>
          </thead>
          <tbody>
            {{ range .Deposits }}
            <tr>
              <td>{{ .BlockNumber }}</td>
              <td>{{ .AssetAmount }}</td>
              <td>{{ .ShareAmount }}</td>
            </tr>
            {{ end }}
          </tbody>
        </table>
      </div>
    </div>

    <div class="box">
      <h2>User Withdrawals</h2>
      <div class="list">
        <table>
          <thead>
            <tr>
              <th>Block Number</th>
              <th>Share Burned</th>
              <th>Asset Recieved</th>
            </tr>
          </thead>
          <tbody>
            {{ range .Deposits }}
            <tr>
              <td>{{ .BlockNumber }}</td>
              <td>{{ .ShareAmount }}</td>
              <td>{{ .AssetAmount }}</td>
            </tr>
            {{ end }}
          </tbody>
        </table>
      </div>
    </div>
     </div>


    </div>

	<div class="heading">
    <img src="https://pon-app.netlify.app/static/media/icon-builder.7fb453f0.svg" alt="icon" style="padding: 10px;">
    Performance
</div>

    <div class="containerFull">
    
    <div class="box">
      <h2>Deposits</h2>
      <div class="chart">
      <canvas id="performanceChart"> </canvas>
      <script>
          var ctx = document.getElementById('performanceChart');
          var performanceChart = new Chart(ctx, {
              data: {
                  labels: {{ .DepositBlocks}},
                  datasets: [{
                    type: 'line',
                    label: 'Deposits',
                    data: {{ .DepositList}},
					
                    borderColor: 'rgba(255, 99, 132, 0.8)',
                    backgroundColor: 'rgba(255, 99, 132, 0.2)',
                    borderDash: [10, 5],
                    yAxisID: 'yLine'
                  }
				  ]
              },
              options: {
                responsive: true,
                scales: {
                    yLine: {
                      type: 'linear',
                      display: true,
                      position: 'left',
                    },
                    yBar: {
                      type: 'linear',
                      display: true,
                      position: 'right',
                      beginAtZero: true,
              
                      grid: {
                        drawOnChartArea: false,
                      },
                    },
                }
              }
          });
      </script>
      </div>
    </div>

</div>

    <div class="containerFull">
    
    <div class="box">
      <h2>Withdrawals</h2>
      <div class="chart">
      <canvas id="depositChart"> </canvas>
      <script>
          var ctx = document.getElementById('depositChart');
          var depositChart = new Chart(ctx, {
              data: {
                  labels: {{ .WithdrawalBlocks}},
                  datasets: [{
                    type: 'line',
                    label: 'Withdrawals',
                    data: {{ .WithdrawList}},
					
                    borderColor: 'rgba(124, 252, 0, 0.8)',
                    backgroundColor: 'rgba(124, 252, 0, 0.2)',
                    borderDash: [10, 5],
                    yAxisID: 'yLine'
                  }
				  ]
              },
              options: {
                responsive: true,
                scales: {
                    yLine: {
                      type: 'linear',
                      display: true,
                      position: 'left',
                    },
                    yAxes: [{
                      display: true,
                      scaleLabel: {
                        display: true,
                        labelString: 'Value'
                      }
                    }],
                    yBar: {
                      type: 'linear',
                      display: true,
                      position: 'right',
                      beginAtZero: true,
              
                      grid: {
                        drawOnChartArea: false,
                      },
                    },
                }
              }
          });
      </script>
      </div>
    </div>

</div>



  <div class="footer">
    Copyright © 2023
  </div>

  
</body>

</html>
      

`)
}

func parseNoDeposits() (*template.Template, error) {
	return template.New("noDeposits").Parse(
		`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Superform APY</title>
  <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
  <style>

    body {
      background-color: black;
      font-family: Arial, sans-serif;
      color: white;
    }

    .heading {
        font-size: 2.25rem;
        line-height: 2.5rem;
        font-weight: 600;
        color: white;
        display: flex;
        padding-top: 40px;
        align-items: center;
        justify-content: center;
    }
    
    .container {
      display: grid;
      grid-template-columns: repeat(2, 1fr);
      gap: 20px;
      margin: 20px;
    }

    .containerFull {
        margin: 20px;
    }
    
    .box {
      border-radius: 10px;
      padding: 20px;
      background-color: #2d2e35;
      box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5);
	  overflow:hidden;
	  overflow-y:scroll;
	  height: 90vh;
	  
    }
    
    .box h2 {
      font-size: 1.5rem;
      margin-top: 0;
      color: #00ed76;
    }
    
    .list {
      margin-top: 10px;
	  overflow:hidden;
	  overflow-y:scroll;
    }
    
    table {
      width: 100%;
      border-collapse: collapse;
      margin-top: 10px;
	  padding-bottom: 50px;
    }
    
    th, td {
      padding: 5px;
	  font-size: 1rem;
      text-align: left;
    }
    
    th {
      border-bottom: 1px solid white;
    }
    
    .chart {
      margin-top: 20px;
      height: 70vh;
      display: flex;
      justify-content: center;
    }

    .finePrint {
        padding-top: 10px;
        display: flex;
        justify-content: center;
        color: #2d2e35;
    }

    .footer {
        margin-top: 80px;
        padding-bottom: 30px;
        padding-top: 30px;
        background-color: #18191a;
        display: flex;
        justify-content: center;
    }

  </style>
</head>
<body>
<div class="heading">
    <img src="https://pon-app.netlify.app/static/media/icon-builder.7fb453f0.svg" alt="icon" style="padding: 10px;">
    No Deposits For User
</div>

  
</body>

</html>
      

`)
}
