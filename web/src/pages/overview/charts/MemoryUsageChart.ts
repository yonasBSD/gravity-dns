import { ChartData } from "chart.js";
import { RolesTsdbApi, TypesAPIMetricsGetOutput, TypesAPIMetricsRole } from "gravity-api";

import { customElement } from "lit/decorators.js";

import { DEFAULT_CONFIG } from "../../../api/Config";
import { groupBy } from "../../../common/utils";
import { getColorFromString } from "../../../elements/charts/Chart";
import { AKChart } from "../../../elements/charts/Chart";

@customElement("gravity-overview-charts-memory-usage")
export class MemoryUsageChart extends AKChart<TypesAPIMetricsGetOutput> {
    apiRequest(): Promise<TypesAPIMetricsGetOutput> {
        return new RolesTsdbApi(DEFAULT_CONFIG).tsdbGetMetrics({
            role: TypesAPIMetricsRole.System,
            category: "memory",
        });
    }

    getChartType(): string {
        return "line";
    }

    getOptions() {
        const opts = super.getOptions();
        opts.scales!.y!.min = 0;
        return opts;
    }

    getChartData(data: TypesAPIMetricsGetOutput): ChartData {
        const chartData: ChartData = {
            datasets: [],
        };
        groupBy(data?.records || [], (record) => record.node).forEach(([node, records]) => {
            const background = getColorFromString(node);
            background.a = 0.3;
            chartData.datasets.push({
                label: node,
                borderColor: getColorFromString(node).toString(),
                backgroundColor: background.toString(),
                spanGaps: true,
                fill: "origin",
                cubicInterpolationMode: "monotone",
                tension: 0.4,
                data: records.map((record) => {
                    return {
                        x: record.time.getTime(),
                        y: Math.round(record.value / 1024 / 1024),
                    };
                }),
            });
        });
        return chartData;
    }
}
