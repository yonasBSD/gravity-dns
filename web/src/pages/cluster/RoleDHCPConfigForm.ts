import { DhcpRoleConfig, RolesDhcpApi } from "gravity-api";

import { TemplateResult, html } from "lit";
import { customElement } from "lit/decorators.js";

import { DEFAULT_CONFIG } from "../../api/Config";
import { first } from "../../common/utils";
import "../../elements/forms/HorizontalFormElement";
import { ModelForm } from "../../elements/forms/ModelForm";

@customElement("gravity-cluster-role-dhcp-config")
export class RoleDHCPConfigForm extends ModelForm<DhcpRoleConfig, string> {
    async loadInstance(): Promise<DhcpRoleConfig> {
        const config = await new RolesDhcpApi(DEFAULT_CONFIG).dhcpGetRoleConfig();
        return config.config;
    }

    getSuccessMessage(): string {
        if (this.instance) {
            return "Successfully updated role config.";
        } else {
            return "Successfully created role config.";
        }
    }

    send = (data: DhcpRoleConfig): Promise<unknown> => {
        return new RolesDhcpApi(DEFAULT_CONFIG).dhcpPutRoleConfig({
            dhcpAPIRoleConfigInput: {
                config: data,
            },
        });
    };

    renderForm(): TemplateResult {
        return html` <ak-form-element-horizontal label="Port" required name="port">
                <input type="number" value="${first(this.instance?.port, 67)}" required />
            </ak-form-element-horizontal>
            <ak-form-element-horizontal
                label="Lease negotiation timeout"
                required
                name="leaseNegotiateTimeout"
                helperText="Time in seconds a client has to acknowledge an IP after it has been offered."
            >
                <input
                    type="number"
                    value="${first(this.instance?.leaseNegotiateTimeout, 30)}"
                    required
                />
            </ak-form-element-horizontal>`;
    }
}
